package auth

import (
	"github.com/alex-305/bookbackend/auth/helpers"
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func SignUp(creds models.Credentials, db *db.DB) (string, error) {

	hashedPassword, err := helpers.HashPassword(creds.Password)

	if err != nil {
		return "", err
	}

	err = helpers.ValidateUsername(creds.Username)

	if err != nil {
		return "", err
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
