package auth

import (
	"github.com/alex-305/bookbackend/auth/helpers"
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func Login(creds models.Credentials, tok string, db *db.DB) (string, error) {
	user, err := db.GetUser(creds.Username)

	if err != nil {
		return "", err
	}

	err = helpers.ComparePassword(creds.Password, user.Password)

	if err != nil {
		return "", err
	}

	err = token.Validate(tok)

	if err != nil {
		tok, err = token.Generate(user.Username)
		if err != nil {
			return "", err
		}
	}
	return tok, nil
}
