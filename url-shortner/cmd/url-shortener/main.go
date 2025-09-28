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
	"url-shortner/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func setupServer() (*http.Server, *gorm.DB) {
	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Init DB
	db := database.New()

	// Init Router
	router := gin.Default()
	routes.SetupRoutes(router, db)

	// Return HTTP Server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	return srv, db.DB
}

func startServer(srv *http.Server) {
	go func() {
		log.Println("Server started on port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}()
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

func markUrlsExpiredPeriodically(db *gorm.DB) {
	go func() {
		for {
			if err := services.MarkExpiredURLs(db); err != nil {
				log.Printf("Error marking URLs as expired: %v", err)
			}
			time.Sleep(1 * time.Minute) // Run every minute
		}
	}()
}

func main() {
	loadEnv()
	srv, db := setupServer()
	markUrlsExpiredPeriodically(db)
	startServer(srv)
	gracefulShutdown(srv)
}
