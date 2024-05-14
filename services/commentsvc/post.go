package commentsvc

import (
	"github.com/alex-305/bookbackend/auth/token"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func Post(r models.Comment, tok models.Token, d *db.DB) (string, error) {
	var err error
	r.Username, err = token.Validate(tok)

	if err != nil {
		return "", err
	}

	replyid, err := d.PostComment(r)

	if err != nil {
		return "", err
	}

	return replyid, nil
}
