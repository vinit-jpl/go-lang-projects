package main

import (
	"fmt"
	"net/http"
)

func main() {
	urlShortener := &UrlShortener{
		urls: make(map[string]string),
	}

	http.HandleFunc("/shorten", urlShortener.HandleShorten)
	http.HandleFunc("/short/", urlShortener.HandleRedirect)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
