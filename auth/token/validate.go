package token

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Validate(token string) (string, error) {
	godotenv.Load()
	secretKey := os.Getenv("SECRET_KEY")

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	return claims.username, nil
}
