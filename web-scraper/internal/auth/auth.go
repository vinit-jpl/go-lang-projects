package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the request headers.
func GetAPIKey(headers http.Header) (string, error) {

	value := headers.Get("Authorization") 

	if value == "" {
		return  "", errors.New("no API key provided")
	}

	values := strings.Split(value, " ")

	if len(values) != 2 {
		return "", errors.New("invalid Authorization header format")

	}

	if values[0] != "ApiKey" {
		return "", errors.New("invalid Authorization header prefix")
	}

	return values[1], nil
}