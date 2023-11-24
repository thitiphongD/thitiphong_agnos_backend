package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	DB_HOST := os.Getenv("POSTGRES_HOST")
	DB_NAME := os.Getenv("POSTGRES_DB")
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PORT := os.Getenv("POSTGRES_PORT")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("Error while opening database connection: %s", err)
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting underlying database object: %s", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	return db
}