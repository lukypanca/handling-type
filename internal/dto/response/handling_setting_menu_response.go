package response

import "database/sql"

type HandlingSettingMenuRepository struct {
	db *sql.DB
}

func NewHandlingSettingMenuRepository(db *sql.DB) *HandlingSettingMenuRepository {
	return &HandlingSettingMenuRepository{db: db}
}

// func (r *HandlingSettingMenuRepository) GetHandlingSetting(

// )
