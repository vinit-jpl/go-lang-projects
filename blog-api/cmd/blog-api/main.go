package main

// Importing necessary packages
import (
	"context" // For context.TODO()
	"fmt"
	"log" // For logging errors and information
	"net/http"
	"os"

	"github.com/joho/godotenv"                      // Loads environment variables from a .env file
	"github.com/vinit-jpl/blog-api/internal/config" // Custom package for configuration and MongoDB connection
	"github.com/vinit-jpl/blog-api/internal/controllers"
	"github.com/vinit-jpl/blog-api/internal/repository"
	"github.com/vinit-jpl/blog-api/internal/routes"
	"github.com/vinit-jpl/blog-api/internal/services"
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

	repo := repository.NewBlogRepository(mongoInstance.DB)
	service := services.NewBlogService(repo)
	postController := controllers.NewPostController(service)

	// http.HandleFunc("POST /post", postController.Create)

	router := http.NewServeMux()
	routes.RegisterPostRoutes(router, postController)
	routes.RegisterViewPostRoutes(router, postController)
	routes.RegisterViewAllPostsRoutes(router, postController)
	routes.RegisterUpdatePostRoutes(router, postController)

	// Start  server here...
	port := os.Getenv("PORT")
	// fmt.Println("port", port)

	fmt.Println("Server started on port,", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("failded to start the server: ", err)
	}

}
