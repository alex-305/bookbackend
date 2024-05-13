package usersvc

import (
	"log"

	"github.com/alex-305/bookbackend/auth/access"
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func PatchDescription(username, description string, tok models.Token, db *db.DB) error {

	userUsername, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = access.HasOwnershipAccess(userUsername, username)

	if err != nil {
		return err
	}

	err = db.UpdateUserDescription(username, description)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}
