package reviewlikesvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Post(reviewid string, tok models.Token, db *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = db.PostReviewLikes(username, reviewid)

	if err != nil {
		return err
	}

	return nil
}
