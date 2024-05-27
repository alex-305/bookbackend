package commentsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Post(c models.Comment, tok models.Token, d *db.DB) (models.CommentID, error) {
	var err error
	c.Username, err = token.Validate(tok)

	if err != nil {
		return "", err
	}

	replyid, err := d.PostComment(c)

	if err != nil {
		return "", err
	}

	return replyid, nil
}
