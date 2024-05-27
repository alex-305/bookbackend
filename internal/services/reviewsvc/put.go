package reviewsvc

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/auth/access"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Put(reviewid models.ReviewID, rev models.Review, tok models.Token, d *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	userID, err := d.GetUserID(username)

	if err != nil {
		return err
	}

	reviewToUpdate, err := d.GetReview(userID, reviewid)

	if err != nil {
		return err
	}

	err = access.HasOwnershipAccess(userID, reviewToUpdate.UserID)

	if err != nil {
		return errors.New("unauthorized")
	}

	err = d.UpdateReview(reviewid, rev)

	if err != nil {
		return err
	}

	return nil
}
