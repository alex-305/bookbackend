package followsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func PostFollow(followed models.Username, tok models.Token, db *db.DB) error {
	follower, err := token.Validate(tok)

	if err != nil {
		return err
	}

	followerID, err := db.GetUserID(follower)

	if err != nil {
		return err
	}

	followedID, err := db.GetUserID(followed)

	if err != nil {
		return err
	}

	err = db.PostFollow(followerID, followedID)

	if err != nil {
		return err
	}

	return nil
}
