package main

import (
	"log"
	"net/http"

	"github.com/rashidbazarbaev/urlshortener/database"
	"github.com/rashidbazarbaev/urlshortener/handler"
	"github.com/rashidbazarbaev/urlshortener/internal"
)

func main() {
	go func() {
		http.HandleFunc("/", handler.RedirectHandler)
		http.ListenAndServe(":8123", nil)
	}()

	go func() {
		if err := database.InitDB(); err != nil {
			log.Fatal(err)
		}
	}()
	internal.Start()
}
