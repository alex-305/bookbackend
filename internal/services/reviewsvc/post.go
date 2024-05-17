package reviewsvc

import (
	"log"

	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Post(rev models.Review, tok models.Token, d *db.DB) (string, error) {
	username, err := token.Validate(tok)

	log.Printf("user:%s", username)

	if err != nil {
		return "", err
	}

	reviewid, err := d.PostReview(username, rev)

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	return reviewid, nil
}
