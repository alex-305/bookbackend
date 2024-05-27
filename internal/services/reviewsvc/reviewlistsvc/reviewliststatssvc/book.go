package reviewliststatssvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetBook(volumeid models.VolumeID, d *db.DB) (models.ReviewListStats, error) {
	rs, err := d.GetBookReviewListStats(volumeid)

	if err != nil {
		return models.ReviewListStats{}, err
	}

	return rs, nil
}
