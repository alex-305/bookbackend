package review

import (
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
)

func PostReview(tok, content string, rating uint8, d *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = d.PostReview(username, content, rating)

	if err != nil {
		return err
	}
	return nil
}
