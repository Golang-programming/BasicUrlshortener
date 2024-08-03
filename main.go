package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalUrl string    `json:"originalUrl"`
	ShortUrl    string    `json:"shortUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

var urlDB = make(map[string]URL)

func generateShortUrl(originalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalUrl))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)

	return hash[:8]
}

func getUrl(id string) string {
	if url, ok := urlDB[id]; ok {
		return url.OriginalUrl
	}

	return ""
}

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

func handleCreateUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "POST" {
		var url URL
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		shortUrl := createUrl(url.OriginalUrl)

		json.NewEncoder(w).Encode(shortUrl)
	}

}

func handleRedirectToUrl(w http.ResponseWriter, r *http.Request) {
	// get id from path
	urlId := r.URL.Path[len("/redirect/"):]

	originalUrl := getUrl(urlId)

	if originalUrl != "" {
		http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
	} else {
		http.NotFound(w, r)
	}
}

func main() {

	// PO method
	http.HandleFunc("/shortner", handleCreateUrl)
	http.HandleFunc("/redirect", handleRedirectToUrl)

	// Starting http server
	fmt.Println("Server running on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
