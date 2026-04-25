package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func postgresConfig(prefix string) DBConfig {
	return DBConfig{
		DSN:     os.Getenv(prefix + "_PG_DSN"),
		MaxOpen: getEnvInt(prefix+"_PG_MAX_OPEN", 10),
		MaxIdle: getEnvInt(prefix+"_PG_MAX_IDLE", 2),
	}
}

func InitPostgres() *Databases {

	return &Databases{
		mufcms: connectPostgres("MUFCMS", postgresConfig("MUFCMS")),
		mufam:  connectPostgres("MUFAM", postgresConfig("MUFAM")),
	}
}

func connectPostgres(name string, cfg DBConfig) *sql.DB {

	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		log.Fatalf("%s DB open error: %v", name, err)
	}

	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetMaxIdleConns(cfg.MaxIdle)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("%s DB ping failed: %v", name, err)
	}

	log.Printf("%s Postgres connected", name)
	return db
}
