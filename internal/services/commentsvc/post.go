package commentsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
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
