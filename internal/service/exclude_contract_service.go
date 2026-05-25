package service

import (
	"context"
	"database/sql"
	"fmt"
	"tipe-handling/internal/enum"
	"tipe-handling/internal/mapper"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/model"
	"tipe-handling/internal/repository/am"
	"tipe-handling/internal/repository/outbox"
)

type ExcludeContractService struct {
	db     *sql.DB
	amrepo *am.ExcludeContractRepository
	// cmsrepo	*cms.ExcludeContractRepository
	outboxrepo *outbox.Repository
}

func NewExcludeContractService(db *sql.DB, amrepo *am.ExcludeContractRepository, outboxrepo *outbox.Repository) *ExcludeContractService {
	return &ExcludeContractService{
		db:         db,
		amrepo:     amrepo,
		outboxrepo: outboxrepo,
	}
}

// func (s *ExcludeContractService) InsertExcludeContractFromExcel(ctx context.Context, rows [][]string) error {

// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	auditInfo := audit.FromContext(ctx)

// 	for i, row := range rows {

// 		if i == 0 || len(row) == 0 {
// 			continue
// 		}

// 		data := mapper.MapExcludeContractToModel(row)

// 		if data.ContractNo == nil {
// 			return fmt.Errorf("row %d: contract_no wajib", i+1)
// 		}

// 		// 1. INSERT DATA
// 		if err := s.amrepo.InsertExcludeContractFromExcel(ctx, tx, data, auditInfo); err != nil {
// 			return fmt.Errorf("row %d error: %w", i+1, err)
// 		}

// 		// 2. CREATE OUTBOX (DALAM LOOP YANG SAMA)
// 		event := outbox.OutboxEvent{
// 			EventType: enum.ExcludeContractCreated,
// 			Payload:   mapper.ToJSON(data),
// 			Status:    outbox.StatusNew,
// 			Retry:     0,
// 		}

// 		if err := s.outboxrepo.Insert(ctx, event); err != nil {
// 			return fmt.Errorf("row %d outbox error: %w", i+1, err)
// 		}
// 	}

// 	return tx.Commit()
// }

func (s *ExcludeContractService) InsertExcludeContractFromExcel(ctx context.Context, rows [][]string) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	auditInfo := audit.FromContext(ctx)

	var items []model.CmsArExcludeContract

	for i, row := range rows {

		if i == 0 || len(row) == 0 {
			continue
		}

		data := mapper.MapExcludeContractToModel(row)

		if data.ContractNo == nil {
			return fmt.Errorf("row %d: contract_no wajib", i+1)
		}

		// 1. insert AM (tetap per row)
		if err := s.amrepo.InsertExcludeContractFromExcel(ctx, tx, data, auditInfo); err != nil {
			return fmt.Errorf("row %d error: %w", i+1, err)
		}

		// 2. collect untuk batch event
		items = append(items, data)
	}

	// commit data dulu
	if err := tx.Commit(); err != nil {
		return err
	}

	// 3. OUTBOX SINGLE EVENT (BATCH)
	payload := struct {
		Items []model.CmsArExcludeContract `json:"items"`
	}{
		Items: items,
	}

	event := outbox.OutboxEvent{
		EventType: enum.ExcludeContractCreated,
		Payload:   mapper.ToJSON(payload),
		Status:    outbox.StatusNew,
		Retry:     0,
	}

	return s.outboxrepo.Insert(ctx, event)
}