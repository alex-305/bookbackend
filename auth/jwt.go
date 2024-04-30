package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Claims struct {
	jwt.StandardClaims
	username string
}

func GenerateJWT(username string) (string, error) {
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

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) error {
	godotenv.Load()
	secretKey := os.Getenv("SECRET_KEY")

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

}
