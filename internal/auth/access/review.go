package access

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Review(username, reviewid string, db *db.DB) error {
	var review models.Review
	review, err := db.GetReview(reviewid)

	if err != nil {
		return err
	}

	return HasOwnershipAccess(username, review.Username)
}
