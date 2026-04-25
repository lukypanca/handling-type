package repository

import (
	"context"
	"database/sql"
	"tipe-handling/internal/model"
)

type HandlingSettingRepository struct {
	CMS *sql.DB
	AM  *sql.DB
}

func NewHandlingSettingRepository(cms *sql.DB, am *sql.DB) *HandlingSettingRepository {
	return &HandlingSettingRepository{
		CMS: cms,
		AM:  am,
	}
}

func (r *HandlingSettingRepository) FindAll(ctx context.Context) ([]model.HandlingSetting, error) {

	query := `
		SELECT HANDLING_SETTING_ID, TIPE_HANDLING, DESC_HANDLING
		FROM MUFCMS.AR_HANDLING_SETTING
	`

	rows, err := r.AM.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.HandlingSetting
	for rows.Next() {
		var m model.HandlingSetting

		if err := rows.Scan(&m.ID, &m.TipeHandling, &m.DescHandling); err != nil {
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
}

func (r *HandlingSettingRepository) Save(ctx context.Context, handling *model.HandlingSetting) (int, error) {

	query := `
		INSERT INTO MUFAM.CMS_AR_HANDLING_SETTING_ (
			DESC_HANDLING,
			TIPE_HANDLING,
			START_OD,
			END_OD
		) VALUES (
			:1, :2, :3, :4
		)
		RETURNING HANDLING_SETTING_ID INTO :5
	`

	var id int

	_, err := r.AM.ExecContext(
		ctx,
		query,
		handling.DescHandling,
		handling.TipeHandling,
		handling.StartOD,
		handling.EndOD,
		sql.Out{Dest: &id},
	)

	if err != nil {
		return 0, err
	}

	return id, nil
}
