package usersvc

import (
	"log"

	"github.com/alex-305/bookbackend/internal/auth/access"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func PatchDescription(username models.Username, description string, tok models.Token, db *db.DB) error {

	userUsername, err := token.Validate(tok)

	if err != nil {
		return err
	}
	useruserID, err := db.GetUserID(userUsername)

	userid, err := db.GetUserID(username)

	user, err := db.GetUser(useruserID, userid)

	err = access.HasOwnershipAccess(useruserID, user.UserID)

	if err != nil {
		return err
	}

	err = db.UpdateUserDescription(userid, description)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}
