package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model         // embeds ID, CreatedAt, UpdatedAt, DeletedAt fields
	OriginalURL string `gorm:"not null" json:"original_url"`
	ShortCode   string `gorm:"uniqueIndex;not null" json:"short_code"`
}
