package services

import (
	"errors"
	"url-shortner/internal/models"

	"gorm.io/gorm"
)

func GetOriginalURL(db *gorm.DB, shortCode string) (string, error) {
	var url models.URL

	err := db.Where("short_code = ?", shortCode).First(&url).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("URL not found")
		}
		return "", err
	}

	// check for expired urls in db

	if url.Expired {
		return "", errors.New("URL has expired")
	}

	return url.OriginalURL, nil
}
