package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"urlshortener/internal/models"
	"urlshortener/internal/storage"
	"urlshortener/pkg/utils"

	"github.com/go-chi/chi/v5"
)

type URLHandler struct {
	storage storage.Storage
	baseURL string
}

func NewURLHandler(storage storage.Storage, baseURL string) *URLHandler {
	return &URLHandler{
		storage: storage,
		baseURL: baseURL,
	}
}

// CreateShortURL handles the creation of a new short URL
func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req models.CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate short code
	shortCode, err := utils.GenerateShortCode()
	if err != nil {
		http.Error(w, "Failed to generate short code", http.StatusInternalServerError)
		return
	}

	// Create URL record
	url := &models.URL{
		ShortCode:    shortCode,
		OriginalURL:  req.URL,
		CreatedAt:    time.Now(),
		ClickCount:   0,
		LastAccessed: time.Time{},
	}

	if err := h.storage.CreateURL(r.Context(), url); err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	// Prepare response
	resp := models.CreateURLResponse{
		ShortURL:    h.baseURL + "/" + shortCode,
		OriginalURL: req.URL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// Redirect handles the redirection to the original URL
func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")
	if shortCode == "" {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	url, err := h.storage.GetURL(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Increment click count
	if err := h.storage.IncrementClickCount(r.Context(), shortCode); err != nil {
		// Log error but continue with redirect
		// TODO: Add proper logging
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}

// GetStats retrieves statistics for a short URL
func (h *URLHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")
	if shortCode == "" {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	stats, err := h.storage.GetURLStats(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
