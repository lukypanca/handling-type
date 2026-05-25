package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"
	"tipe-handling/internal/enum"
	audit "tipe-handling/internal/metadata"
	"tipe-handling/internal/model"
	"tipe-handling/internal/repository/cms"
	"tipe-handling/internal/repository/outbox"
)

type CmsSyncExcContWorker struct {
	db *sql.DB
	outboxrepo	*outbox.Repository
	cmsrepo		*cms.ExcludeContractRepository
	interval	time.Duration
}

func NewCmsSyncExcContWorker(
	db *sql.DB,
	outboxrepo *outbox.Repository,
	cmsrepo *cms.ExcludeContractRepository,
) *CmsSyncExcContWorker {
	return &CmsSyncExcContWorker{
		db:	db,
		outboxrepo: outboxrepo,
		cmsrepo: cmsrepo,
		interval:	2 * time.Second,
	}
}

func (w *CmsSyncExcContWorker) Start(ctx context.Context) {

	log.Println("CMS Sync Worker Execlude Contract started")

	for {
		select {
		case <-ctx.Done():
			log.Println("CMS Sync Worker stopped")
			return
		default:
			w.process(ctx)
		}

		time.Sleep(w.interval)
	}
}

func (w *CmsSyncExcContWorker) process(ctx context.Context) {

	events, err := w.outboxrepo.FindPending(ctx, enum.ExcludeContractCreated, 50)
	if err != nil {
		log.Println("failed fetch outbox exclude contract: ", err)
		return
	}

	for _, event := range events {

		err := w.handleEvent(ctx, event)
		if err != nil {
			log.Println("failed sync exclude contract event: ", event.ID, err)

			_ = w.outboxrepo.MarkFailed(ctx, event.ID)
			continue
		}

		_ = w.outboxrepo.MarkProcessed(ctx, event.ID)
	}
}

func (w *CmsSyncExcContWorker) handleEvent(ctx context.Context, event outbox.OutboxEvent) error {

	switch event.EventType {

	case enum.ExcludeContractCreated:
		return w.syncExcludeCreated(ctx, event.Payload)
	
	default:
		log.Println("unknown event exclude contract: ", event.EventType)
		return nil
	}
}

// func (w *CmsSyncExcContWorker) syncExcludeCreated(ctx context.Context, payload string) error {

// 	var data model.CmsArExcludeContract

// 	if err := json.Unmarshal([]byte(payload), &data); err != nil {
// 		return err
// 	}

// 	tx, err := w.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	auditInfo := audit.FromContext(ctx)

// 	return w.cmsrepo.InsertExcludeContractFromExcel(ctx, tx, data, auditInfo)
// }

func (w *CmsSyncExcContWorker) syncExcludeCreated(ctx context.Context, payload string) error {

	var req struct {
		Items []model.CmsArExcludeContract `json:"items"`
	}

	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		return err
	}

	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	auditInfo := audit.FromContext(ctx)

	for _, data := range req.Items {

		if err := w.cmsrepo.InsertExcludeContractFromExcel(ctx, tx, data, auditInfo); err != nil {
			return err
		}
	}

	return tx.Commit()
}