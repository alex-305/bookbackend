package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetFollowing(tok models.Token, o models.ReviewSortOptions, db *db.DB) ([]models.Review, error) {
	username, err := token.Validate(tok)

	if err != nil {
		return []models.Review{}, err
	}

	return db.GetFollowingReviewList(username, o)
}
