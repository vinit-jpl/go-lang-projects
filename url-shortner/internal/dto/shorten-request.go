package dto

// data transfer object for shortening URL request

type ShortenRequest struct {
	URL              string `json:"url" binding:"required"`
	CustomCode       string `json:"custom_code"`
	ExpiresInMinutes *int   `json:"expires_in_minutes"`
	Expired          bool   `json:"expired"`
}
