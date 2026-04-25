package config

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	// _ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"
)

type Databases struct {
	MUFCMS *sql.DB
	MUFAM  *sql.DB
}

type DBConfig struct {
	DSN     string
	MaxOpen int
	MaxIdle int
}

func getEnvInt(key string, def int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return def
	}
	return val
}

func loadConfig(prefix string) DBConfig {
	return DBConfig{
		DSN: os.Getenv(prefix + "_DSN"),
		MaxOpen: getEnvInt(prefix + "_MAX_OPEN", 10),
		MaxIdle: getEnvInt(prefix + "_MAX_IDLE", 2),
	}
}

func connectDB(name string, cfg DBConfig) *sql.DB {
	db, err := sql.Open("oracle", cfg.DSN)
	if err != nil {
		log.Fatalf("%s DB open error: %v", name, err)
	}
	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdle) * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatalf("%s DB ping failed: %v", name, err)
	}

	log.Printf("%s DB connected successfully\n", name)
	return db
}

func InitDatabases() *Databases {

	return &Databases{
		MUFCMS: connectDB("MUFCMS", loadConfig("MUFCMS")),
		MUFAM:  connectDB("MUFAM", loadConfig("MUFAM")),
	}
}
