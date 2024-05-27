package usersvc

import (
	"github.com/alex-305/bookbackend/internal/auth/access"
	"github.com/alex-305/bookbackend/internal/auth/helpers"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func PatchPassword(userToUpdate models.Username, password string, tok models.Token, db *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	userID, err := db.GetUserID(username)

	if err != nil {
		return err
	}

	userToUpdateID, err := db.GetUserID(userToUpdate)

	if err != nil {
		return err
	}

	err = access.HasOwnershipAccess(userID, userToUpdateID)

	if err != nil {
		return err
	}

	hashedPassword, err := helpers.HashPassword(password)

	if err != nil {
		return err
	}

	err = db.UpdateUserPassword(userID, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}
