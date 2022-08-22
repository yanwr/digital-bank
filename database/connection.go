package database

import (
	"fmt"
	"log"
	"os"
	"yanwr/digital-bank/env"
	"yanwr/digital-bank/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func LoadConnectionDB() {
	var err error

	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", os.Getenv(env.DB_HOST), os.Getenv(env.DB_USER), os.Getenv(env.DB_PASS), os.Getenv(env.DB_NAME), os.Getenv(env.DB_TZ))
	dbConnection, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv(env.DB_AUTO_MIGRATE) == "true" {
		loadMigrates()
	}

	log.Print("Connected with DB")
}

func loadMigrates() {
	dbConnection.AutoMigrate(&models.Account{})
	dbConnection.AutoMigrate(&models.Transfer{})
}

func GetConnectionDB() *gorm.DB {
	return dbConnection
}
