package config

import (
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	DSN     string
	MaxOpen int
	MaxIdle int
}

func getEnvInt(key string, def int) int {
	val := os.Getenv(key)

	if val == "" {
		return def
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}

	return i
}

func Init() DB {

	log.Println("DB_TYPE RAW =", "["+os.Getenv("DB_TYPE")+"]")

	switch os.Getenv("DB_TYPE") {

	case "oracle":
		return InitOracle()

	case "postgres":
		return InitPostgres()

	default:
		panic("DB_TYPE not set")
	}
}
