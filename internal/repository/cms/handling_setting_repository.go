package cms

import (
	"context"
	"database/sql"
	"tipe-handling/internal/enum"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/model"
)

type HandlingSettingRepository struct {
	db *sql.DB
}

func NewHandlingSettingRepository(db *sql.DB) *HandlingSettingRepository {
	return &HandlingSettingRepository{db: db}
}

func (r *HandlingSettingRepository) SaveHandlingSetting(
	ctx context.Context, tx *sql.Tx, handling *model.HandlingSetting, audit audit.Info) (int, error) {

	query := enum.InsertHandlingSettingCms

	var id int

	_, err := tx.ExecContext(
		ctx,
		query,
		handling.DescHandling,
		handling.TipeHandling,
		handling.StartOD,
		handling.EndOD,
		handling.Status,
		handling.IsActive,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
		sql.Out{Dest: &id},
	)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *HandlingSettingRepository) SaveHandlingBranch(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingBranch, audit audit.Info) error {
	query := enum.InsertBranchQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.KodeCabang,
		handling.NamaCabang,
		handling.KodeArea,
		handling.NamaArea,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingObject(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingObject, audit audit.Info) error {
	query := enum.InsertObjectQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.ObjectCode,
		handling.ObjectGroup,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingTipeNasabah(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingTipeNasabah, audit audit.Info) error {
	query := enum.InsertTipeNasabahQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.TipeNasabahCode,
		handling.TipeNasabahDesc,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingCollScoring(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingCollScoring, audit audit.Info) error {
	query := enum.InsertCollScoringQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.CollScoringCode,
		handling.CollScoringDesc,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingPaymentType(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingPaymentType, audit audit.Info) error {
	query := enum.InsertPaymentTypeQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.PaymentTypeCode,
		handling.PaymentTypeDesc,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingTipePembiayaan(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingTipePembiayaan, audit audit.Info) error {
	query := enum.InsertTipePembiayaanQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.TipePembiayaanCode,
		handling.TipePembiayaanDesc,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingSkemaPembiayaan(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingSkemaPembiayaan, audit audit.Info) error {
	query := enum.InsertSkemaPembiayaanQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.SkemaPembiayaanCode,
		handling.SkemaPembiayaanDesc,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingPenggolonganProduct(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingPenggolonganProduct, audit audit.Info) error {
	query := enum.InsertGolonganProductQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.PenggolonganProductCode,
		handling.PenggolonganProductDesc,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSettingRepository) SaveHandlingBankPendanaan(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingBankPendanaan, audit audit.Info) error {
	query := enum.InsertBankPendanaanQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.BankCode,
		handling.BankDesc,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}
