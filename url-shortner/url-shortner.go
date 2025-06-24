package main

import (
	"fmt"
	"net/http"
)

type UrlShortener struct {
	urls map[string]string
}

func (us *UrlShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalUrl := r.FormValue("url")

	if originalUrl == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	shortKey := GenerateShortKey()

	us.urls[shortKey] = originalUrl

	// Return the shortened URL as a response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Shortened URL: http://localhost:8080/short/" + shortKey))

}

func (us *UrlShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url path:", r.URL)
	shortKey := r.URL.Path[len("/short/"):]
	originalUrl, ok := us.urls[shortKey]
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusFound)
}
