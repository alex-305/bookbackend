package commentlistsvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetReview(reviewid string, o models.CommentSortOptions, db *db.DB) ([]models.Comment, error) {
	return db.GetReviewComments(reviewid, o)
}
