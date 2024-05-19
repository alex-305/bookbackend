package db

import "github.com/alex-305/bookbackend/internal/models"

func (db DB) GetBookReviewStats(volumeID string) models.ReviewListStats {
	query := `SELECT COUNT(*) AS reviewcount, AVG(rating) AS avgrating FROM reviews WHERE volumeid = $1`
	return db.getReviewListStats(query, volumeID)
}

func (db DB) GetUserReviewStats(username string) models.ReviewListStats {
	query := `SELECT COUNT(*), AVG(rating) FROM reviews WHERE username = $1`
	return db.getReviewListStats(query, username)
}

func (db DB) getReviewListStats(query, param string) models.ReviewListStats {
	var stats models.ReviewListStats
	_ = db.QueryRow(query, param).Scan(&stats.ReviewCount, &stats.AvgRating)
	return stats
}
