package reviewliststatssvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetUserStats(username models.Username, db *db.DB) (models.ReviewListStats, error) {
	userID, err := db.GetUserID(username)

	if err != nil {
		return models.ReviewListStats{}, err
	}

	rs, err := db.GetUserReviewListStats(userID)

	if err != nil {
		return models.ReviewListStats{}, err
	}

	return rs, nil
}
