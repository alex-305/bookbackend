package commentsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(commentid models.CommentID, tok models.Token, d *db.DB) (models.Comment, error) {
	username, err := token.Validate(tok)
	if err != nil {
		return models.Comment{}, err
	}

	userid, err := d.GetUserID(username)

	if err != nil {
		return models.Comment{}, err
	}

	comment, err := d.GetComment(userid, commentid)

	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}
