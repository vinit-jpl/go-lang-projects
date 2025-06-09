package main

// Importing necessary packages
import (
	"context" // For context.TODO()
	"log"     // For logging errors and information

	"github.com/joho/godotenv"                      // Loads environment variables from a .env file
	"github.com/vinit-jpl/blog-api/internal/config" // Custom package for configuration and MongoDB connection
)

func main() {
	// Load .env file into environment variables
	err := godotenv.Load() // Loads variables from .env into the process environment
	if err != nil {
		// If .env file is missing or loading fails, log a warning (not fatal)
		log.Println("No .env file found or error loading .env:", err)
	}

	mongoInstance, err := config.ConnectMongo()
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	defer mongoInstance.Client.Disconnect(context.TODO())

	log.Println("Connected to MongoDB:", mongoInstance.DB.Name())

	// Start your server here...
}
