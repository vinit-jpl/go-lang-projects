package handlers

import (
	"net/http"
	"url-shortner/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Pinghandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// handler for shorteningn a URL

func ShortenURLHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// req holds the incoming data for a URL shortening request.
		// It uses the ShortenRequest struct defined in the services package.
		var req services.ShortenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		shortCode, err := services.ShortenURL(db, req.URL, req.CustomCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"short_code": c.Request.Host + "/" + shortCode,
		})
	}
}
