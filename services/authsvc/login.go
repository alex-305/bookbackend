package authsvc

import (
	"github.com/alex-305/bookbackend/auth/helpers"
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func Login(creds models.Credentials, tok models.Token, db *db.DB) (models.Token, error) {
	dbcreds, err := db.GetCredentials(creds.Username)

	if err != nil {
		return "", err
	}

	err = helpers.ComparePassword(creds.Password, dbcreds.Password)

	if err != nil {
		return "", err
	}

	_, err = token.Validate(tok)

	if err != nil {
		tok, err = token.Generate(dbcreds.Username)
		if err != nil {
			return "", err
		}
	}
	return tok, nil
}
