package db

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/models/tables"
)

func (db DB) PostCommentLikes(userID models.UserID, commentid models.CommentID) error {
	tx := db.Begin()

	newLike := tables.User_likes_comment{
		UserID:    userID,
		CommentID: commentid,
	}

	result := tx.Create(&newLike)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	result = queries.ChangeCommentLikeCount(commentid, tx, 1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()

	return nil
}

func (db DB) DeleteCommentLikes(username models.UserID, commentid models.CommentID) error {
	tx := db.Begin()
	result := tx.Delete(&tables.User_likes_comment{}, "username=? AND commentid=?", username, commentid)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("could not delete from user_likes_comment")
	}

	result = queries.ChangeCommentLikeCount(commentid, tx, 1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()
	return nil
}
