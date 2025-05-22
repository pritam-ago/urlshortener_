package storage

import (
	"context"
	"encoding/json"
	"time"

	"urlshortener/internal/models"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(redisURL string) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	// Test connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisStorage{client: client}, nil
}

func (s *RedisStorage) CreateURL(ctx context.Context, url *models.URL) error {
	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	return s.client.Set(ctx, url.ShortCode, data, 0).Err()
}

func (s *RedisStorage) GetURL(ctx context.Context, shortCode string) (*models.URL, error) {
	data, err := s.client.Get(ctx, shortCode).Bytes()
	if err != nil {
		return nil, err
	}

	var url models.URL
	if err := json.Unmarshal(data, &url); err != nil {
		return nil, err
	}

	return &url, nil
}

func (s *RedisStorage) IncrementClickCount(ctx context.Context, shortCode string) error {
	url, err := s.GetURL(ctx, shortCode)
	if err != nil {
		return err
	}

	url.ClickCount++
	url.LastAccessed = time.Now()

	return s.CreateURL(ctx, url)
}

func (s *RedisStorage) GetURLStats(ctx context.Context, shortCode string) (*models.URLStats, error) {
	url, err := s.GetURL(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return &models.URLStats{
		ShortCode:    url.ShortCode,
		ClickCount:   url.ClickCount,
		LastAccessed: url.LastAccessed,
		CreatedAt:    url.CreatedAt,
	}, nil
}
