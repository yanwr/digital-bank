package config

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	ENV             = "ENV"
	DB_HOST         = "DB_HOST"
	DB_USER         = "DB_USER"
	DB_PASS         = "DB_PASS"
	DB_NAME         = "DB_NAME"
	DB_TZ           = "DB_TZ"
	DB_TYPE         = "DB_TYPE"
	DB_AUTO_MIGRATE = "DB_AUTO_MIGRATE"

	SERVER_PORT = "SERVER_PORT"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}
