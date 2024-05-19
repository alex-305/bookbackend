package commentsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/access"
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Delete(commentid string, tok models.Token, db *db.DB) error {
	username, err := token.Validate(tok)

	if err != nil {
		return err
	}

	err = access.HasOwnershipAccess(username, commentid)

	if err != nil {
		return err
	}

	err = db.DeleteComment(commentid)

	if err != nil {
		return err
	}

	return nil

}
