package commentsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/access"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Delete(commentid models.CommentID, tok models.Token, db *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	userID, err := db.GetUserID(username)

	if err != nil {
		return err
	}

	comment, err := db.GetComment(userID, commentid)

	err = access.HasOwnershipAccess(userID, comment.UserID)

	if err != nil {
		return err
	}

	err = db.DeleteComment(commentid)

	if err != nil {
		return err
	}

	return nil

}
