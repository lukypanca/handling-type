package service

import (
	"context"
	"database/sql"
	"log"
	"tipe-handling/internal/dto/request"
	"tipe-handling/internal/dto/response"
	"tipe-handling/internal/enum"
	"tipe-handling/internal/mapper"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/model"
	"tipe-handling/internal/repository/am"
	"tipe-handling/internal/repository/cms"
	"tipe-handling/internal/repository/outbox"
)

type HandlingSettingService struct {
	db      *sql.DB
	amrepo  *am.HandlingSettingRepository
	cmsRepo *cms.HandlingSettingRepository
	outboxRepo	*outbox.Repository
}

func NewHandlingSettingService(db *sql.DB, amrepo *am.HandlingSettingRepository, cmsrepo *cms.HandlingSettingRepository, outboxrepo *outbox.Repository) *HandlingSettingService {
	return &HandlingSettingService{
		db:      db,
		amrepo:  amrepo,
		cmsRepo: cmsrepo,
		outboxRepo: outboxrepo,
	}
}

func (s *HandlingSettingService) GetAll(ctx context.Context) ([]model.HandlingSetting, error) {
	return s.amrepo.FindAll(ctx)
}

func (s *HandlingSettingService) Create(
	ctx context.Context,
	req *request.CreateHandlingSettingRequest,
) (*response.CreateHandlingSettingResponse, error) {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	data := mapper.ToHandlingSettingModel(req)
	auditInfo := audit.FromContext(ctx)

	// insert tabel CMS_AR_HANDLING_SETTING_
	id, err := s.amrepo.SaveHandlingSetting(ctx, tx, data, auditInfo)
	if err != nil {
		return nil, err
	}

	data.ID = id

	// insert tabel cms_ar_handling_branch
	for _, r := range req.Branches {
		model := mapper.ToHandlingBranchModel(r, id)
		err := s.amrepo.SaveHandlingBranch(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_HANDLING_OBJECT_GROUP
	for _, r := range req.ObjectGroups {
		model := mapper.ToHandlingObjectModel(r, id)
		err := s.amrepo.SaveHandlingObject(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_HANDLING_TIPE_NASABAH
	for _, r := range req.TipeNasabah {
		model := mapper.ToHandlingTipeNasabah(r, id)
		err := s.amrepo.SaveHandlingTipeNasabah(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_HANDLING_COLL_SCORING
	for _, r := range req.CollScoring {
		model := mapper.ToHandlingCollScoring(r, id)
		err := s.amrepo.SaveHandlingCollScoring(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_HANDLING_PAYMENT_TYPE
	for _, r := range req.PaymentTypes {
		model := mapper.ToHandlingPaymentType(r, id)
		err := s.amrepo.SaveHandlingPaymentType(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}
	// insert tabel CMS_AR_HANDLING_T_PEMBIAYAAN
	for _, r := range req.TipePembiayaan {
		model := mapper.ToHandlingTipePembiayaan(r, id)
		err := s.amrepo.SaveHandlingTipePembiayaan(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}
	// insert tabel CMS_AR_HANDLING_S_PEMBIAYAAN
	for _, r := range req.SkemaPembiayaan {
		model := mapper.ToHandlingSkemaPembiayaan(r, id)
		err := s.amrepo.SaveHandlingSkemaPembiayaan(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}
	// insert tabel CMS_AR_HANDLING_GOL_PRODUK
	for _, r := range req.GolProduk {
		model := mapper.ToHandlingPenggolonganProduct(r, id)
		err := s.amrepo.SaveHandlingPenggolonganProduct(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}
	// insert tabel CMS_AR_HANDLING_BANK_PENDANAAN
	for _, r := range req.BankPendanaan {
		model := mapper.ToHandlingBankPendanaan(r, id)
		err := s.amrepo.SaveHandlingBankPendanaan(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// commit jika semua save data sukses
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// =========================
	// 5. OUTBOX EVENT (ASYNC CMS SYNC)
	// =========================
	event := outbox.OutboxEvent{
		EventType: enum.HandlingCreated,
		Payload:   mapper.ToJSON(req), // atau struct result
		Status:    outbox.StatusNew,
		Retry:     0,
	}

	err = s.outboxRepo.Insert(ctx, event)
	if err != nil {
		log.Println("failed insert outbox:", err)
		return nil, err
	}

	return &response.CreateHandlingSettingResponse{
		ID: id,
	}, nil
}
