package reviewsvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(reviewid string, db *db.DB) (models.Review, error) {
	return db.GetReview(reviewid)
}
