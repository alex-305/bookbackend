package token

import (
	"log"
	"os"

	"github.com/alex-305/bookbackend/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Validate(token models.Token) (string, error) {
	godotenv.Load()
	secretKey := os.Getenv("SECRET_KEY")

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(models.String(token), claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	log.Printf("USERNAME FROM TOKEN:%s", claims.Username)

	return claims.Username, nil
}
