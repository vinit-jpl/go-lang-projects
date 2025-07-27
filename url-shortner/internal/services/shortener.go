package services

import (
	"url-shortner/internal/models"
	"url-shortner/internal/utils"

	"gorm.io/gorm"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func ShortenURL(db *gorm.DB, originalURL string) (string, error) {
	var shortCode string

	//  Generate a unique short code
	for {
		shortCode = utils.GenerateShortCode(6)
		// Define a variable to hold the result (if found)
		var exists models.URL

		// Search the database for any row where short_code == generated shortCode
		result := db.Where("short_code = ?", shortCode).First(&exists)

		// Check if the record was not found
		if result.Error == gorm.ErrRecordNotFound {
			// The shortCode is unique — safe to use
			break
		}
	}

	url := models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	if err := db.Create(&url).Error; err != nil {
		return "", err
	}

	return shortCode, nil
}
