package handlers

import (
	"net/http"
	"url-shortner/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RedirectHandler returns a handler for redirecting short codes to original URLs
func RedirectHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the short code from the path
		code := c.Param("code")

		// 2. Fetch original URL from the database
		originalURL, err := services.GetOriginalURL(db, code)
		if err != nil {
			// If not found, return 404
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}

		// 3. Redirect to the original URL
		c.Redirect(http.StatusFound, originalURL)
	}
}
