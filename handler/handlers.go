package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/rashidbazarbaev/urlshortener/database"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]

	var originalURL string

	err := database.DB.QueryRow(context.Background(),
		"SELECT original_url FROM short_urls WHERE code = $1",
		code,
	).Scan(&originalURL)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("ngrok-skip-browser-warning", "true")

	http.Redirect(w, r, originalURL, http.StatusFound)
}
