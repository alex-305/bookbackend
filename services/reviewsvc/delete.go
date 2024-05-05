package reviewsvc

import (
	"github.com/alex-305/bookbackend/auth/access"
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
)

func Delete(tok, reviewid string, db *db.DB) error {
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
