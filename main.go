package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// URL represents a URL record
type URL struct {
	ID          string    `json:"id"`
	OriginalUrl string    `json:"originalUrl"`
	ShortUrl    string    `json:"shortUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

// In-memory database to store URLs
var urlDB = make(map[string]URL)

// generateShortUrl creates a shortened version of the original URL
func generateShortUrl(originalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalUrl))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)

	return hash[:8] // Use the first 8 characters of the hash
}

// getUrl retrieves the original URL by its shortened version
func getUrl(id string) string {
	if url, ok := urlDB[id]; ok {
		return url.OriginalUrl
	}

	return ""
}

// createUrl generates a short URL and stores it in the in-memory database
func createUrl(originalUrl string) (shortUrl string) {
	shortUrl = generateShortUrl(originalUrl)

	urlDB[shortUrl] = URL{
		ID:          shortUrl,
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
		CreatedAt:   time.Now(),
	}

	return
}

// handleCreateUrl handles requests to shorten URLs
func handleCreateUrl(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Url string `json:"url"`
	}

	if r.Method == "POST" {
		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		shortUrl := createUrl(body.Url)

		json.NewEncoder(w).Encode(shortUrl)
	}
}

// handleRedirectToUrl handles redirection from shortened URL to the original URL
func handleRedirectToUrl(w http.ResponseWriter, r *http.Request) {
	urlId := r.URL.Path[len("/redirect/"):]

	originalUrl := getUrl(urlId)

	if originalUrl != "" {
		http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
	} else {
		http.NotFound(w, r)
	}
}

// main function to set up routes and start the server
func main() {
	http.HandleFunc("/shortner", handleCreateUrl)
	http.HandleFunc("/redirect/", handleRedirectToUrl)

	fmt.Println("Server running on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
