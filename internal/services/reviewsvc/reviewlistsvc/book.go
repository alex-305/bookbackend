package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetBook(volumeid string, o models.ReviewSortOptions, db *db.DB) ([]models.Review, error) {
	return db.GetBookReviewList(volumeid, o)
}
