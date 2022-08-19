package database

import (
	"fmt"
	"log"
	"os"
	"yanwr/digital-bank/config"
	"yanwr/digital-bank/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func LoadConnectionDB() {
	var err error

	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", os.Getenv(config.DB_HOST), os.Getenv(config.DB_USER), os.Getenv(config.DB_PASS), os.Getenv(config.DB_NAME), os.Getenv(config.DB_TZ))
	dbConnection, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv(config.DB_AUTO_MIGRATE) == "true" {
		dbConnection.AutoMigrate(&models.Account{})
		dbConnection.AutoMigrate(&models.Transfer{})
	}

	log.Print("Connected with DB")
}

func GetConnectionDB() *gorm.DB {
	return dbConnection
}
