package review

import (
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func Get(reviewid string, db *db.DB) (models.Review, error) {
	return db.GetReview(reviewid)
}
