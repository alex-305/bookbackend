package reviewstatssvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(reviewid string, d *db.DB) models.ReviewStats {
	return d.GetReviewStats(reviewid)
}
