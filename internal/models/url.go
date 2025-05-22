package models

import "time"

type URL struct {
	ShortCode    string    `json:"short_code"`
	OriginalURL  string    `json:"original_url"`
	CreatedAt    time.Time `json:"created_at"`
	ClickCount   int64     `json:"click_count"`
	LastAccessed time.Time `json:"last_accessed,omitempty"`
}

type CreateURLRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type CreateURLResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type URLStats struct {
	ShortCode    string    `json:"short_code"`
	ClickCount   int64     `json:"click_count"`
	LastAccessed time.Time `json:"last_accessed"`
	CreatedAt    time.Time `json:"created_at"`
}
