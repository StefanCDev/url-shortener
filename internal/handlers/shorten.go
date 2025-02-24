package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/db"
	"url-shortener/internal/utils"
)

// Request structure
type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

// Response structure
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// ShortenHandler handles URL shortening requests
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.LongURL == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique short code
	shortCode := utils.GenerateShortCode()

	// Save the URL in DynamoDB
	err = db.SaveURL(shortCode, req.LongURL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	// Return response
	response := ShortenResponse{ShortURL: fmt.Sprintf("https://yourdomain.com/%s", shortCode)}
	json.NewEncoder(w).Encode(response)
}
