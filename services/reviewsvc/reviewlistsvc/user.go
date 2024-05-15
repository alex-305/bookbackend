package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func GetUser(username string, o models.ReviewSortOptions, db *db.DB) ([]models.Review, error) {
	return db.GetUserReviewList(username, o)
}
