package token

import (
	"errors"
	"os"
	"time"

	"github.com/alex-305/bookbackend/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Claims struct {
	jwt.StandardClaims
	Username string
}

func Generate(username string) (models.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("failure to create jwt token")
	}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * (24 * 7)).Unix()

	godotenv.Load()
	SecretKey := os.Getenv("SECRET_KEY")

	if SecretKey == "" {
		return "", errors.New("secret key could not be retrieved")
	}

	tokStr, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}
	return models.NewToken(tokStr), nil
}
