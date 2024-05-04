package review

import (
	"log"

	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func Post(tok string, rev models.Review, d *db.DB) (string, error) {
	username, err := token.Validate(tok)

	log.Printf("user:%s", username)

	if err != nil {
		return "", err
	}

	reviewid, err := d.PostReview(username, rev)

	if err != nil {
		return "", err
	}
	return reviewid, nil
}
