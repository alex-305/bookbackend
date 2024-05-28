package db

import (
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) PostReview(userID models.UserID, rev models.NewReview) (models.ReviewID, error) {

	err := db.Table("users").Create(&rev).Error

	if err != nil {
		return "", err
	}

	return rev.ReviewID, nil
}

func (db *DB) DeleteReview(reviewid models.ReviewID) error {
	var reviews []models.Review
	err := db.Delete(&reviews, reviewid).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetReview(userid models.UserID, reviewid models.ReviewID) (models.Review, error) {
	var rev models.Review
	err := db.Table("reviews r").
		Select(`r.userid, r.volumeid, r.reviewid, r.content, r.rating, r.post_date, r.likecount,
	CASE WHEN ulr.userid IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked`).
		Joins(queries.JoinReviewLikes(), userid).
		//Joins(queries.JoinForUsername(queries.ReviewTableName())).
		Where("r.reviewid = ?", reviewid).
		Scan(&rev).
		Error

	if err != nil {
		return models.Review{}, err
	}

	return rev, nil
}

func (db *DB) UpdateReview(reviewid models.ReviewID, r models.Review) error {
	updates := map[string]interface{}{
		"Content": r.Content,
		"Rating":  r.Rating,
	}
	err := db.Model(&models.Review{}).Where("reviewid = ?", reviewid).Updates(updates).Error
	if err != nil {
		return err
	}
	return err
}
