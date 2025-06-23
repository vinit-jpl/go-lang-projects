package utilities

import (
	"os"
	"path/filepath"
)

func CreateFile(filename string) (*os.File, error) {

	// create an uploads directory if it doesn't exist

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	// building the file path and creating the file

	dst, err := os.Create(filepath.Join("uploads", filename))

	if err != nil {
		return nil, err
	}

	return dst, nil
}
