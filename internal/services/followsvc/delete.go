package followsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func DeleteFollow(followed string, tok models.Token, d *db.DB) error {
	follower, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = d.DeleteFollow(follower, followed)

	if err != nil {
		return err
	}

	return nil
}
