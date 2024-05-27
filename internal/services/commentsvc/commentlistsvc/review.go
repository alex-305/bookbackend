package commentlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetReview(reviewid models.ReviewID, tok models.Token, o models.CommentSortOptions, db *db.DB) ([]models.Comment, error) {
	username, err := token.Validate(tok)
	var userID models.UserID
	if err == nil {
		userID, _ = db.GetUserID(username)
	}

	return db.GetReviewComments(userID, reviewid, o)
}
