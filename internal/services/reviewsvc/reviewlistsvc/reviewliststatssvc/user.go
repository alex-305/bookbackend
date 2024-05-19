package reviewliststatssvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetUser(username string, d *db.DB) (models.ReviewListStats, error) {
	return d.GetUserReviewStats(username), nil
}
