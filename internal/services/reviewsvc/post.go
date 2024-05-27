package reviewsvc

import (
	"log"

	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Post(rev models.Review, tok models.Token, d *db.DB) (models.ReviewID, error) {
	username, err := token.Validate(tok)

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	userID, err := d.GetUserID(username)

	reviewid, err := d.PostReview(userID, rev)

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	return reviewid, nil
}
