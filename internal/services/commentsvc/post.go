package commentsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Post(com models.NewComment, tok models.Token, d *db.DB) (models.CommentID, error) {
	var err error
	com.Username, err = token.Validate(tok)

	if err != nil {
		return "", err
	}

	replyid, err := d.PostComment(com)

	if err != nil {
		return "", err
	}

	return replyid, nil
}
