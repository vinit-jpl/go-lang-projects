package services

import (
	"fmt"
	"time"
	"url-shortner/internal/models"
	"url-shortner/internal/utils"

	"gorm.io/gorm"
)

// service heleprs for shortening URL

// check if custom is already taken or not

func isCustomCodeAvailabe(db *gorm.DB, customCode string) error {
	var exists models.URL

	result := db.Where("short_code = ?", customCode).First(&exists)

	if result.Error == nil {
		// Found an existing record -> custom code is already taken
		return fmt.Errorf("custom code '%s' is already in use", customCode)
	}

	if result.Error != gorm.ErrRecordNotFound {
		// Some other DB error
		return result.Error
	}

	return nil
}

// generate a unique random code
func generateUniqueCode(db *gorm.DB, length int) (string, error) {
	for {
		code := utils.GenerateShortCode(length)
		var exists models.URL
		result := db.Where("short_code = ?", code).First(&exists) // checking if the radmon code already exists
		if result.Error == gorm.ErrRecordNotFound {
			return code, nil
		}
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return "", result.Error
		}
	}
}

// calculate expiry
func calculateExpiry(minutes *int) *time.Time {
	if minutes == nil {
		return nil
	}
	t := time.Now().Add(time.Duration(*minutes) * time.Minute)
	return &t
}

// MarkExpiredURLs updates URLs whose expiry time has passed
func MarkExpiredURLs(db *gorm.DB) error {
	now := time.Now()

	// Update all URLs where ExpiresAt is set, in the past, and not yet marked as expired
	result := db.Model(&models.URL{}).
		Where("expires_at IS NOT NULL AND expires_at <= ? AND expired = ?", now, false).
		Update("expired", true)

	return result.Error
}
