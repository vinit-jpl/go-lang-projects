package main

import (
	"fmt"
	"go-file-uploader/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", handlers.FileUploadHandler)

	// Serve static files (HTML, JS, CSS)
	fs := http.FileServer(http.Dir("internal/web"))
	http.Handle("/", fs)

	fmt.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
