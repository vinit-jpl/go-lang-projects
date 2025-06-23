package handlers

import (
	"fmt"
	"go-file-uploader/internal/utilities"
	"io"
	"net/http"
	"strings"
)

// validating file type before saving, for now only supporting image upload
func isValidFileType(file []byte) bool {
	fileType := http.DetectContentType(file)
	return strings.HasPrefix(fileType, "image/")
}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	/*
		limiting the file size to 10 mb
		left shift operator. Shift the bits of 10, 20 times to the left
		effectively it is 10 * 2^20.
	*/
	r.ParseMultipartForm(10 << 20)

	// retrieve the file from form handler

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	if !isValidFileType(fileBytes) {
		http.Error(w, "Invalid file type", http.StatusUnsupportedMediaType)
		return
	}

	// saving the file locally

	dst, err := utilities.CreateFile(handler.Filename)

	if err != nil {
		http.Error(w, "Error in saving the file", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// copy the uploaded file to the destination file

	if _, err := dst.Write(fileBytes); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Uploaded file: %s\n", handler.Filename)
	fmt.Fprintf(w, "File Size: %.2f KB\n", float64(handler.Size)/(1024))
	// fmt.Fprintf(w, "MIME Header: %v\n", handler.Header)
	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}
