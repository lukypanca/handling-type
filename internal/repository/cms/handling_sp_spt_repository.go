package cms

import (
	"context"
	"database/sql"
	"tipe-handling/internal/enum"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/model"
)

type HandlingSpSptRepository struct {
	db *sql.DB
}

func NewHandlingSpSptRepository(db *sql.DB) *HandlingSpSptRepository {
	return &HandlingSpSptRepository{db: db}
}

func (r *HandlingSpSptRepository) SaveHandlingSpSpt(
	ctx context.Context, tx *sql.Tx, handling *model.HandlingSpSpt, audit audit.Info) (int, error) {

	query := enum.InsertHandlingSpSptCms

	var id int

	_, err := tx.ExecContext(
		ctx,
		query,
		handling.DescHandling,
		handling.TipeHandling,
		handling.TipeHandlingId,
		handling.StartOD,
		handling.EndOD,
		handling.FlagRod,
		handling.Status,
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

func (r *HandlingSpSptRepository) SaveHandlingBranchSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingBranch, audit audit.Info) error {
	query := enum.InsertSpSptBranchQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.KodeCabang,
		handling.NamaCabang,
		handling.KodeArea,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSpSptRepository) SaveHandlingObjectSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingObject, audit audit.Info) error {
	query := enum.InsertSpSptObjectQueryCms
	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		handling.ObjectGroup,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}

func (r *HandlingSpSptRepository) SaveHandlingTipeNasabahSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingTipeNasabah, audit audit.Info) error {
	query := enum.InsertSpSptTipeNasabahQueryCms
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

func (r *HandlingSpSptRepository) SaveHandlingCollScoringSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingCollScoring, audit audit.Info) error {
	query := enum.InsertSpSptCollScoringQueryCms
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

func (r *HandlingSpSptRepository) SaveHandlingPaymentTypeSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingPaymentType, audit audit.Info) error {
	query := enum.InsertSpSptPaymentTypeQueryCms
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

func (r *HandlingSpSptRepository) SaveHandlingTipePembiayaanSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingTipePembiayaan, audit audit.Info) error {
	query := enum.InsertSpSptTipePembiayaanQueryCms
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

func (r *HandlingSpSptRepository) SaveHandlingSkemaPembiayaanSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingSkemaPembiayaan, audit audit.Info) error {
	query := enum.InsertSpSptSkemaPembiayaanQueryCms
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

func (r *HandlingSpSptRepository) SaveHandlingPenggolonganProductSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingPenggolonganProduct, audit audit.Info) error {
	query := enum.InsertSpSptGolonganProductQueryCms
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

func (r *HandlingSpSptRepository) SaveHandlingBankPendanaanSpSpt(
	ctx context.Context, tx *sql.Tx, id int, handling *model.HandlingBankPendanaan, audit audit.Info) error {
	query := enum.InsertSpSptBankPendanaanQueryCms
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
