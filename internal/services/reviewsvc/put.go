package reviewsvc

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/auth/access"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Put(reviewid string, rev models.Review, tok models.Token, d *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	reviewToUpdate, err := d.GetReview(reviewid)

	if err != nil {
		return err
	}

	err = access.HasOwnershipAccess(username, reviewToUpdate.Username)

	if err != nil {
		return errors.New("unauthorized")
	}

	err = d.UpdateReview(reviewid, rev)

	if err != nil {
		return err
	}

	return nil
}
