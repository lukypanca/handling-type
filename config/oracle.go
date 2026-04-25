package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/sijms/go-ora/v2"
)

func oracleConfig(prefix string) DBConfig {
	return DBConfig{
		DSN:     os.Getenv(prefix + "_DSN"),
		MaxOpen: getEnvInt(prefix+"_MAX_OPEN", 10),
		MaxIdle: getEnvInt(prefix+"_MAX_IDLE", 2),
	}
}

func InitOracle() *Databases {

	return &Databases{
		mufcms: connectOracle("MUFCMS", oracleConfig("MUFCMS")),
		mufam:  connectOracle("MUFAM", oracleConfig("MUFAM")),
	}
}

func connectOracle(name string, cfg DBConfig) *sql.DB {

	db, err := sql.Open("oracle", cfg.DSN)
	if err != nil {
		log.Fatalf("%s DB open error: %v", name, err)
	}

	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdle) * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatalf("%s DB ping failed: %v", name, err)
	}

	log.Printf("%s Oracle connected", name)
	return db
}
