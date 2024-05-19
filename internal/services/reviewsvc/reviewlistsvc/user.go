package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetUser(username string, o models.ReviewSortOptions, db *db.DB) (models.ReviewList, error) {
	return db.GetUserReviewList(username, o)
}
