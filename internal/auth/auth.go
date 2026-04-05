package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	vals := strings.Split(headers.Get("Authorization"), " ")
	if len(vals) != 2 || vals[0] != "APIKey" {
		return "", errors.New("invalid authorization header")
	}
	return vals[1], nil
}
