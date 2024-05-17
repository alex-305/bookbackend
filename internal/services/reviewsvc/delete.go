package reviewsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/access"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Delete(reviewid string, tok models.Token, db *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = access.Review(username, reviewid, db)

	if err != nil {
		return err
	}

	err = db.DeleteReview(reviewid)

	if err != nil {
		return err
	}

	return nil

}
