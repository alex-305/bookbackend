package db

import (
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db DB) GetBookReviewListStats(volumeID models.VolumeID) (models.ReviewListStats, error) {
	var rls models.ReviewListStats
	err := db.Table("? AS inner_query", db.Table("reviews")).
		Select("AVG(rating) AS avguserrating, COUNT(*) AS reviewcount").
		Where("volumeid = ?", volumeID).
		Group("userid").
		Select("SUM(reviewcount) AS reviewcount, AVG(avguserrating) as avgrating").
		Scan(&rls).
		Error

	if err != nil {
		return models.ReviewListStats{}, err
	}

	return rls, nil
}

func (db DB) GetUserReviewListStats(userid models.UserID) (models.ReviewListStats, error) {
	var rls models.ReviewListStats
	err := db.Table(queries.FromReview()).
		Select("COUNT(*) AS reviewcount, AVG(rating) AS avgrating").
		//Joins(queries.JoinForUserID(queries.ReviewTableName()), userid).
		Where("r.userid = ?", userid).
		Scan(&rls).
		Error

	if err != nil {
		return models.ReviewListStats{}, err
	}
	return rls, nil
}
