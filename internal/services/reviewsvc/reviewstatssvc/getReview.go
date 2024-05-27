package reviewstatssvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetReviewStats(reviewid models.ReviewID, d *db.DB) (models.ReviewStats, error) {

	rs, err := d.GetReviewStats(reviewid)

	if err != nil {
		return models.ReviewStats{}, err
	}
	return rs, nil
}
