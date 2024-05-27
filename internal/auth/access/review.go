package access

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Review(userID models.UserID, reviewid models.ReviewID, db *db.DB) error {
	var review models.Review
	review, err := db.GetReview(userID, reviewid)

	if err != nil {
		return err
	}

	return HasOwnershipAccess(userID, review.UserID)
}
