package main

import (
	"log"
	"net/http"

	"urlshortener/internal/config"
	"urlshortener/internal/handlers"
	"urlshortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize configuration
	cfg := config.New()

	// Initialize Redis storage
	redisStorage, err := storage.NewRedisStorage(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Initialize URL handler
	urlHandler := handlers.NewURLHandler(redisStorage, cfg.BaseURL)

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("URL Shortener Service"))
	})

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/shorten", urlHandler.CreateShortURL)
		r.Get("/stats/{shortCode}", urlHandler.GetStats)
	})

	// Redirect route
	r.Get("/{shortCode}", urlHandler.Redirect)

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatal(err)
	}
}
