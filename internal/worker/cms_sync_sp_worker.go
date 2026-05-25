package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"
	"tipe-handling/internal/dto/request"
	"tipe-handling/internal/enum"
	"tipe-handling/internal/mapper"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/repository/cms"
	"tipe-handling/internal/repository/outbox"
)

type CmsSyncSpWorker struct {
	db         *sql.DB
	outboxRepo *outbox.Repository
	cmsRepo    *cms.HandlingSpSptRepository
	interval   time.Duration
}

func NewCmsSyncSpWorker(
	db *sql.DB,
	outboxRepo *outbox.Repository,
	cmsRepo *cms.HandlingSpSptRepository,
) *CmsSyncSpWorker {

	return &CmsSyncSpWorker{
		db:         db,
		outboxRepo: outboxRepo,
		cmsRepo:    cmsRepo,
		interval:   2 * time.Second,
	}
}

func (w *CmsSyncSpWorker) Start(ctx context.Context) {

	log.Println("🚀 CMS Sync SP Worker started")

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 CMS Sync SP Worker stopped")
			return
		default:
			w.process(ctx)
		}

		time.Sleep(w.interval)
	}
}

func (w *CmsSyncSpWorker) process(ctx context.Context) {

	events, err := w.outboxRepo.FindPending(ctx, enum.SpCreated, 50)
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

func (w *CmsSyncSpWorker) handleEvent(ctx context.Context, event outbox.OutboxEvent) error {

	switch event.EventType {

	case enum.SpCreated:
		return w.syncSpCreated(ctx, event.Payload)

	default:
		log.Println("unknown event:", event.EventType)
		return nil
	}
}

func (w *CmsSyncSpWorker) syncSpCreated(ctx context.Context, payload string) error {

	// 1. decode payload
	var req request.CreateHandlingSpSptRequest
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

	data := mapper.ToHandlingSpSptModel(&req)
	auditInfo := audit.FromContext(ctx)

	// insert tabel AR_HANDLING_SETTING_
	id, err := w.cmsRepo.SaveHandlingSpSpt(ctx, tx, data, auditInfo)
	if err != nil {
		return err
	}

	// insert tabel AR_HANDLING_BRANCH
	for _, r := range req.Branches {
		model := mapper.ToHandlingBranchModel(r, id)
		if err := w.cmsRepo.SaveHandlingBranchSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_OBJECT_GROUP
	for _, r := range req.ObjectGroups {
		model := mapper.ToHandlingObjectModel(r, id)
		if err := w.cmsRepo.SaveHandlingObjectSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_TIPE_NASABAH
	for _, r := range req.TipeNasabah {
		model := mapper.ToHandlingTipeNasabah(r, id)
		if err := w.cmsRepo.SaveHandlingTipeNasabahSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_COLL_SCORING
	for _, r := range req.CollScoring {
		model := mapper.ToHandlingCollScoring(r, id)
		if err := w.cmsRepo.SaveHandlingCollScoringSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_PAYMENT_TYPE
	for _, r := range req.PaymentTypes {
		model := mapper.ToHandlingPaymentType(r, id)
		if err := w.cmsRepo.SaveHandlingPaymentTypeSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_T_PEMBIAYAAN
	for _, r := range req.TipePembiayaan {
		model := mapper.ToHandlingTipePembiayaan(r, id)
		if err := w.cmsRepo.SaveHandlingTipePembiayaanSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_S_PEMBIAYAAN
	for _, r := range req.SkemaPembiayaan {
		model := mapper.ToHandlingSkemaPembiayaan(r, id)
		if err := w.cmsRepo.SaveHandlingSkemaPembiayaanSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_GOL_PRODUK
	for _, r := range req.GolProduk {
		model := mapper.ToHandlingPenggolonganProduct(r, id)
		if err := w.cmsRepo.SaveHandlingPenggolonganProductSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// insert tabel AR_HANDLING_BANK_PENDANAAN
	for _, r := range req.BankPendanaan {
		model := mapper.ToHandlingBankPendanaan(r, id)
		if err := w.cmsRepo.SaveHandlingBankPendanaanSpSpt(ctx, tx, id, model, auditInfo); err != nil {
			return err
		}
	}

	// commit
	return tx.Commit()
}
