package commentlikesvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Delete(commentID models.CommentID, tok models.Token, db *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	userID, err := db.GetUserID(username)

	if err != nil {
		return err
	}

	err = db.DeleteCommentLikes(userID, commentID)

	if err != nil {
		return err
	}

	return nil
}
