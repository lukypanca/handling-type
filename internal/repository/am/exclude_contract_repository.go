package am

import (
	"context"
	"database/sql"
	"tipe-handling/internal/enum"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/model"
)

type ExcludeContractRepository struct {
	db *sql.DB
}

func NewExcludeContractRepository(db *sql.DB) *ExcludeContractRepository {
	return &ExcludeContractRepository{db: db}
}

func (r *ExcludeContractRepository) InsertExcludeContractFromExcel(ctx context.Context, tx *sql.Tx, data model.CmsArExcludeContract, audit audit.Info) error {

	query := enum.InsertCmsArExcludeContract

	_, err := tx.ExecContext(
		ctx,
		query,
		data.ContractNo,
		data.NamaNasabah,
		data.ObjectGroup,
		data.TipeNasabah,
		data.CollScoring,
		data.PaymentType,
		data.TipePembiayaan,
		data.SkemaPembiayaan,
		data.PenggolonganProduct,
		data.BankPendanaan,
		data.MarketingProgram,
		data.TipeHandling,
		data.StatusExclude,
		1,
		audit.Now,
		audit.User,
		audit.Now,
		audit.User,
	)

	return err
}
