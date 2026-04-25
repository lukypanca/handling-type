package config

import "database/sql"

type DB interface {
	MUFCMS() *sql.DB
	MUFAM() *sql.DB
	Close()
}
