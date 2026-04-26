package outbox

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(ctx context.Context, event OutboxEvent) error {

	query := `
		INSERT INTO MUFAM.OUTBOX_EVENT (
			EVENT_TYPE,
			PAYLOAD,
			STATUS,
			RETRY,
			CREATED_AT,
			UPDATED_AT
		) VALUES (
			:1, :2, :3, :4, SYSDATE, SYSDATE
		)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		event.EventType,
		event.Payload,
		string(event.Status),
		event.Retry,
	)

	return err
}

func (r *Repository) FindPending(ctx context.Context, limit int) ([]OutboxEvent, error) {

	query := `
		SELECT ID, EVENT_TYPE, PAYLOAD, STATUS, RETRY, CREATED_AT, UPDATED_AT
		FROM MUFAM.OUTBOX_EVENT
		WHERE STATUS = 'NEW'
		ORDER BY CREATED_AT
		FETCH FIRST :1 ROWS ONLY
	`

	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []OutboxEvent

	for rows.Next() {
		var e OutboxEvent

		if err := rows.Scan(
			&e.ID,
			&e.EventType,
			&e.Payload,
			&e.Status,
			&e.Retry,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, e)
	}

	return result, nil
}

func (r *Repository) MarkProcessed(ctx context.Context, id int) error {

	query := `
		UPDATE MUFAM.OUTBOX_EVENT
		SET STATUS = 'PROCESSED',
		    UPDATED_AT = SYSDATE
		WHERE ID = :1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *Repository) MarkFailed(ctx context.Context, id int) error {

	query := `
		UPDATE MUFAM.OUTBOX_EVENT
		SET STATUS = 'FAILED',
		    RETRY = RETRY + 1,
		    UPDATED_AT = SYSDATE
		WHERE ID = :1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
