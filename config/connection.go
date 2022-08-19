package config

import (
	"fmt"
	"log"
	"os"
	"yanwr/digital-bank/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var dbConnection *gorm.DB
	var err error

	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", os.Getenv(DB_HOST), os.Getenv(DB_USER), os.Getenv(DB_PASS), os.Getenv(DB_NAME), os.Getenv(DB_TZ))
	fmt.Println("yAN HERE" + DNS)
	dbConnection, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv(DB_AUTO_MIGRATE) == "true" {
		dbConnection.AutoMigrate(&models.Account{})
		dbConnection.AutoMigrate(&models.Transfer{})
	}

	return dbConnection
}
