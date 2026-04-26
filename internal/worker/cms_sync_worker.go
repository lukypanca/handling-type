package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"
	"tipe-handling/internal/dto/request"
	"tipe-handling/internal/mapper"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/repository/cms"
	"tipe-handling/internal/repository/outbox"
)

type CmsSyncWorker struct {
	db         *sql.DB
	outboxRepo *outbox.Repository
	cmsRepo    *cms.HandlingSettingRepository
	interval   time.Duration
}

func NewCmsSyncWorker(
	db *sql.DB,
	outboxRepo *outbox.Repository,
	cmsRepo *cms.HandlingSettingRepository,
) *CmsSyncWorker {

	return &CmsSyncWorker{
		db:         db,
		outboxRepo: outboxRepo,
		cmsRepo:    cmsRepo,
		interval:   2 * time.Second,
	}
}

func (w *CmsSyncWorker) Start(ctx context.Context) {

	log.Println("🚀 CMS Sync Worker started")

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 CMS Sync Worker stopped")
			return
		default:
			w.process(ctx)
		}

		time.Sleep(w.interval)
	}
}

func (w *CmsSyncWorker) process(ctx context.Context) {

	events, err := w.outboxRepo.FindPending(ctx, 50)
	if err != nil {
		log.Println("failed fetch outbox:", err)
		return
	}

	for _, event := range events {

		err := w.handleEvent(ctx, event)
		if err != nil {
			log.Println("failed sync event:", event.ID, err)

			_ = w.outboxRepo.MarkFailed(ctx, event.ID)
			continue
		}

		_ = w.outboxRepo.MarkProcessed(ctx, event.ID)
	}
}

func (w *CmsSyncWorker) handleEvent(ctx context.Context, event outbox.OutboxEvent) error {

	switch event.EventType {

	case "HANDLING_CREATED":
		return w.syncHandlingCreated(ctx, event.Payload)

	default:
		log.Println("unknown event:", event.EventType)
		return nil
	}
}

func (w *CmsSyncWorker) syncHandlingCreated(ctx context.Context, payload string) error {

	// 1. decode payload
	var req request.CreateHandlingSettingRequest
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		return err
	}

	// 2. begin transaction CMS
	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	data := mapper.ToHandlingSettingModel(&req)
	auditInfo := audit.FromContext(ctx)

	// insert tabel AR_HANDLING_SETTING_
	id, err := w.cmsRepo.SaveHandlingSetting(ctx, tx, data, auditInfo)
	if err != nil {
		return err
	}

	// insert tabel AR_HANDLING_BRANCH
	for _, r := range req.Branches {
		model := mapper.ToHandlingBranchModel(r, id)
		if err := w.cmsRepo.SaveHandlingBranch(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_OBJECT_GROUP
	for _, r := range req.ObjectGroups {
		model := mapper.ToHandlingObjectModel(r, id)
		if err := w.cmsRepo.SaveHandlingObject(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}	

	// insert tabel AR_HANDLING_TIPE_NASABAH
	for _, r := range req.TipeNasabah {
		model := mapper.ToHandlingTipeNasabah(r, id)
		if err := w.cmsRepo.SaveHandlingTipeNasabah(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}
	
	// insert tabel AR_HANDLING_COLL_SCORING
	for _, r := range req.CollScoring {
		model := mapper.ToHandlingCollScoring(r, id)
		if err := w.cmsRepo.SaveHandlingCollScoring(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}
	
	// insert tabel AR_HANDLING_PAYMENT_TYPE
	for _, r := range req.PaymentTypes {
		model := mapper.ToHandlingPaymentType(r, id)
		if err := w.cmsRepo.SaveHandlingPaymentType(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}	

	// insert tabel AR_HANDLING_T_PEMBIAYAAN
	for _, r := range req.TipePembiayaan {
		model := mapper.ToHandlingTipePembiayaan(r, id)
		if err := w.cmsRepo.SaveHandlingTipePembiayaan(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}
	
	// insert tabel AR_HANDLING_S_PEMBIAYAAN
	for _, r := range req.SkemaPembiayaan {
		model := mapper.ToHandlingSkemaPembiayaan(r, id)
		if err := w.cmsRepo.SaveHandlingSkemaPembiayaan(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_GOL_PRODUK
	for _, r := range req.GolProduk {
		model := mapper.ToHandlingPenggolonganProduct(r, id)
		if err := w.cmsRepo.SaveHandlingPenggolonganProduct(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_BANK_PENDANAAN
	for _, r := range req.BankPendanaan {
		model := mapper.ToHandlingBankPendanaan(r, id)
		if err := w.cmsRepo.SaveHandlingBankPendanaan(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// commit
	return tx.Commit()
}
