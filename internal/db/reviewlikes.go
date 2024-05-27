package db

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/models/tables"
)

func (db DB) PostReviewLikes(userID models.UserID, reviewid models.ReviewID) error {
	tx := db.Begin()

	newLike := tables.User_likes_review{
		UserID:   userID,
		ReviewID: reviewid,
	}

	result := tx.Create(&newLike)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	result = queries.ChangeReviewLikeCount(reviewid, tx, 1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()

	return nil
}

func (db DB) DeleteReviewLikes(userID models.UserID, reviewid models.ReviewID) error {
	tx := db.Begin()
	result := tx.Delete(&tables.User_likes_comment{}, "userid=? AND reviewid=?", userID, reviewid)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("could not delete from user_likes_review")
	}

	result = queries.ChangeReviewLikeCount(reviewid, tx, -1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()
	return nil
}
