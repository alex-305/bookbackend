package db

import (
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetReviewStats(reviewid models.ReviewID) (models.ReviewStats, error) {
	var rs models.ReviewStats
	err := db.Table(queries.FromComment()).
		Select("COUNT(*) AS commentcount").
		Where("reviewid = ?", reviewid).
		Scan(&rs).
		Error
	if err != nil {
		return models.ReviewStats{}, err
	}
	return rs, nil
}
