package helpers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/alex-305/bookbackend/internal/models"
)

func GetToken(r *http.Request) (models.Token, error) {
	var tok models.Token
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	tok = models.NewToken(parts[1])

	return tok, nil
}
