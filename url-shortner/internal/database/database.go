package database

import (
	"fmt"
	"log"
	"os"
	"url-shortner/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func New() *Database {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",

		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	// fmt.Println("Connecting to database with DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connection established successfully")

	if err := db.AutoMigrate(&models.URL{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	myDatabase := Database{DB: db}

	return &myDatabase
}
