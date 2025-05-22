package config

import (
	"os"
)

type Config struct {
	RedisURL    string
	ServerPort  string
	BaseURL     string
	Environment string
}

func New() *Config {
	return &Config{
		RedisURL:    getEnv("REDIS_URL", "localhost:6379"),
		ServerPort:  getEnv("PORT", "8080"),
		BaseURL:     getEnv("BASE_URL", "http://localhost:8080"),
		Environment: getEnv("ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
