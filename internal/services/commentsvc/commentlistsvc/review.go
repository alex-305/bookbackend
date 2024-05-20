package commentlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetReview(reviewid string, tok models.Token, o models.CommentSortOptions, db *db.DB) ([]models.Comment, error) {
	username, _ := token.Validate(tok)
	return db.GetReviewComments(username, reviewid, o)
}
