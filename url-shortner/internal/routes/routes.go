package routes

import (
	"url-shortner/internal/database"
	"url-shortner/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *database.Database) {
	// Example route
	router.GET("/ping", handlers.Pinghandler)

	// routes/routes.go
	v1 := router.Group("/v1")
	{
		v1.POST("/shorten", handlers.ShortenURLHandler(db.DB))
	}
	router.GET("/:code", handlers.RedirectHandler(db.DB))

}
