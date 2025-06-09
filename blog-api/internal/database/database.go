package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func InitDB() {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	pingErr := DB.Ping() // check if the database connection is working

	if pingErr != nil {
		log.Fatal("Error pinging the database:", pingErr)
	}

	log.Println("Database connection established successfully")

}
