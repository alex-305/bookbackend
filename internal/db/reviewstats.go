package db

import "github.com/alex-305/bookbackend/internal/models"

func (db *DB) GetReviewStats(reviewid string) models.ReviewStats {
	query := `SELECT COUNT(*) AS commentcount FROM comments WHERE reviewid=$1`
	return db.getReviewStats(query, reviewid)
}

func (db DB) getReviewStats(query, param string) models.ReviewStats {
	var stats models.ReviewStats
	_ = db.QueryRow(query, param).Scan(&stats.CommentCount)
	return stats
}
