package authsvc

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/auth/helpers"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func SignUp(creds models.Credentials, db *db.DB) (models.Token, error) {

	hashedPassword, err := helpers.HashPassword(creds.Password)

	if err != nil {
		return "", err
	}

	ok := helpers.ValidateUsername(creds.Username)

	if !ok {
		return "", errors.New("invalid credentials")
	}

	ok = helpers.ValidateEmail(creds.Email)

	if !ok {
		return "", errors.New("invalid credentials")
	}

	creds.Password = hashedPassword

	err = db.CreateUser(creds)

	if err != nil {
		return "", err
	}

	token, err := token.Generate(creds.Username)

	if err != nil {
		return "", err
	}

	return token, nil

}
