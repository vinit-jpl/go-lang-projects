package routes

import (
	"url-shortner/internal/database"
	"url-shortner/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *database.Database) {
	// Example route
	router.GET("/ping", handlers.Pinghandler)

	// You can pass db.DB (the *gorm.DB) to your handlers here
}
