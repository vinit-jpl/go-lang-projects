package services

import (
	"url-shortner/internal/dto"
	"url-shortner/internal/models"

	"gorm.io/gorm"
)

func ShortenURL(db *gorm.DB, req dto.ShortenRequest) (string, error) {
	var shortCode string

	// if user has provided a custom code, check its availability

	if req.CustomCode != "" {
		err := isCustomCodeAvailabe(db, req.CustomCode)

		if err != nil {
			return "", err
		}

		shortCode = req.CustomCode
	} else {
		// generate a unique random code of length 6
		code, err := generateUniqueCode(db, 6)

		if err != nil {
			return "", err
		}

		shortCode = code
	}

	// create a new URL record in the database
	url := models.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
		ExpiresAt:   calculateExpiry(req.ExpiresInMinutes),
		Expired:     false,
	}

	// save the URL record to the database
	if err := db.Create(&url).Error; err != nil {
		return "", err
	}

	return shortCode, nil

}
