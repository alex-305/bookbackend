package usersvc

import (
	"github.com/alex-305/bookbackend/auth/access"
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
)

func PatchDescription(tok, username, description string, db *db.DB) error {

	userUsername, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = access.HasOwnershipAccess(userUsername, username)

	if err != nil {
		return err
	}

	err = db.UpdateUserDescription(username, description)

	return nil
}
