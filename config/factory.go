package config

import (
	"log"
	"os"
	"strconv"
	"strings"
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

func GetBool(key string, def bool) bool {
	val := strings.ToLower(strings.TrimSpace(os.Getenv(key)))

	if val == "" {
		return def
	}

	switch val {
	case "true", "1", "yes", "y", "on":
		return true
	case "false", "0", "no", "n", "off":
		return false
	default:
		return def
	}
}