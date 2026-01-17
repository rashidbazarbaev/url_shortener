package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func isValidUrl(link string) bool {
	u, err := url.Parse(link)
	if err != nil || u.Scheme == "" || u.Host == "" {
		log.Println("invalid link")
		return false
	}
	return true
}

func IsReachable(link string) bool {
	if !isValidUrl(link) {
		return false
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(link)
	if err != nil {
		fmt.Printf("Error reaching URL %s: %v\n", link, err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 300 {
		return true
	}

	fmt.Printf("Url %s returned status code %d\n", link, resp.StatusCode)
	return false
}
