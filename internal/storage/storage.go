package storage

import (
	"context"
	"urlshortener/internal/models"
)

type Storage interface {
	// URL operations
	CreateURL(ctx context.Context, url *models.URL) error
	GetURL(ctx context.Context, shortCode string) (*models.URL, error)
	IncrementClickCount(ctx context.Context, shortCode string) error
	GetURLStats(ctx context.Context, shortCode string) (*models.URLStats, error)
}
