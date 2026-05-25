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
	"tipe-handling/internal/repository/am"
	"tipe-handling/internal/repository/cms"
	"tipe-handling/internal/repository/outbox"
)

type HandlingSpSptService struct {
	db         *sql.DB
	amrepo     *am.HandlingSpSptRepository
	cmsrepo    *cms.HandlingSpSptRepository
	outboxrepo *outbox.Repository
}

func NewHandlingSpSptService(db *sql.DB, amrepo *am.HandlingSpSptRepository, cmsrepo *cms.HandlingSpSptRepository, outboxrepo *outbox.Repository) *HandlingSpSptService {
	return &HandlingSpSptService{
		db:         db,
		amrepo:     amrepo,
		cmsrepo:    cmsrepo,
		outboxrepo: outboxrepo,
	}
}

func (s *HandlingSpSptService) Create(
	ctx context.Context,
	req *request.CreateHandlingSpSptRequest,
) (*response.CreateHandlingSpSptResponse, error) {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	data := mapper.ToHandlingSpSptModel(req)
	auditInfo := audit.FromContext(ctx)

	// insert tabel CMS_AR_H_SP_SPT_
	id, err := s.amrepo.SaveHandlingSpSpt(ctx, tx, data, auditInfo)
	if err != nil {
		return nil, err
	}

	data.ID = id

	// insert tabel cms_ar_h_sp_spt_branch
	for _, r := range req.Branches {
		model := mapper.ToHandlingBranchModel(r, id)
		err := s.amrepo.SaveHandlingBranchSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_H_SP_SPT_OBJECT_GROUP
	for _, r := range req.ObjectGroups {
		model := mapper.ToHandlingObjectModel(r, id)
		err := s.amrepo.SaveHandlingObjectSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_H_SP_SPT_TIPE_NASABAH
	for _, r := range req.TipeNasabah {
		model := mapper.ToHandlingTipeNasabah(r, id)
		err := s.amrepo.SaveHandlingTipeNasabahSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_H_SP_SPT_COLL_SCORING
	for _, r := range req.CollScoring {
		model := mapper.ToHandlingCollScoring(r, id)
		err := s.amrepo.SaveHandlingCollScoringSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_H_SP_SPT_PAYMENT_TYPE
	for _, r := range req.PaymentTypes {
		model := mapper.ToHandlingPaymentType(r, id)
		err := s.amrepo.SaveHandlingPaymentTypeSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_H_SP_SPT_T_PEMBIAYAAN
	for _, r := range req.TipePembiayaan {
		model := mapper.ToHandlingTipePembiayaan(r, id)
		err := s.amrepo.SaveHandlingTipePembiayaanSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_H_SP_SPT_S_PEMBIAYAAN
	for _, r := range req.SkemaPembiayaan {
		model := mapper.ToHandlingSkemaPembiayaan(r, id)
		err := s.amrepo.SaveHandlingSkemaPembiayaanSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// insert tabel CMS_AR_H_SP_SPT_GOL_PRODUK
	for _, r := range req.GolProduk {
		model := mapper.ToHandlingPenggolonganProduct(r, id)
		err := s.amrepo.SaveHandlingPenggolonganProductSpSpt(ctx, tx, id, model, auditInfo)
		if err != nil {
			return nil, err
		}
	}

	// inser tabel CMS_AR_H_SP_SPT_BANK_PENDANAAN
	for _, r := range req.BankPendanaan {
		model := mapper.ToHandlingBankPendanaan(r, id)
		err := s.amrepo.SaveHandlingBankPendanaanSpSpt(ctx, tx, id, model, auditInfo)
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
		EventType: enum.SpCreated,
		Payload:   mapper.ToJSON(req), // atau struct result
		Status:    outbox.StatusNew,
		Retry:     0,
	}

	err = s.outboxrepo.Insert(ctx, event)
	if err != nil {
		log.Println("failed insert outbox:", err)
		return nil, err
	}

	return &response.CreateHandlingSpSptResponse{
		ID: id,
	}, nil
}
