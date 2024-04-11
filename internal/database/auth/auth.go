package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an Api key from the headers of an http request
// ex: Authorization: ApiKey {api key here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentification info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}
	return vals[1], nil
}
