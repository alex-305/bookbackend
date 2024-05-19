package reviewliststatssvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetBook(volumeid string, d *db.DB) (models.ReviewListStats, error) {
	return d.GetBookReviewStats(volumeid), nil
}
