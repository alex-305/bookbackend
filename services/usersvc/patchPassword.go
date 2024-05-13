package usersvc

import (
	"github.com/alex-305/bookbackend/auth/access"
	"github.com/alex-305/bookbackend/auth/helpers"
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func PatchPassword(userToUpdate, password string, tok models.Token, db *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = access.HasOwnershipAccess(username, userToUpdate)

	if err != nil {
		return err
	}

	hashedPassword, err := helpers.HashPassword(password)

	if err != nil {
		return err
	}

	err = db.UpdateUserPassword(username, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}
