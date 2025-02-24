package handlers

import (
	"fmt"
	"net/http"
	"url-shortener/internal/db"
)

// RedirectHandler handles redirections from short URLs
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // Extract short code from URL

	// Fetch the original URL
	longURL, err := db.GetURL(shortCode)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// Redirect user to the original URL
	http.Redirect(w, r, longURL, http.StatusFound)
}
