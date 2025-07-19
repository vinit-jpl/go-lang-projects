package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortner/internal/database"
	"url-shortner/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variales

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialze gorm database connection

	db := database.New()

	// Initialize gin router

	router := gin.Default()

	// Register routes
	routes.SetupRoutes(router, db)

	// Create http server

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// start the server in go routine
	go func() {
		log.Println("Server started on port:", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// shutdown context with timeout

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancle()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting gracefully")

}
