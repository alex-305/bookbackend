package authsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/helpers"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Login(creds models.Credentials, tok models.Token, db *db.DB) (models.Token, error) {
	userID, err := db.GetUserID(creds.Username)

	if err != nil {
		return "", err
	}

	dbcreds, err := db.GetCredentials(userID)

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
