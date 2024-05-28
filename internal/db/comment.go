package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) PostComment(com models.NewComment) (models.CommentID, error) {

	err := db.Table("comments").Create(&com).Error

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	return com.CommentID, nil
}

func (db *DB) DeleteComment(commentid models.CommentID) error {
	var comments []models.Comment
	result := db.Delete(&comments, commentid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DB) GetComment(userID models.UserID, commentID models.CommentID) (models.Comment, error) {
	var comment models.Comment

	result :=
		db.Table(queries.FromComment()).
			Select(queries.SelectComment()).
			Joins(queries.JoinCommentLikes(), userID).
			//Joins(queries.JoinForUsername(queries.CommentTableName())).
			Where("c.commentid = ?", commentID).
			Scan(&comment)
	if result.Error != nil {
		return models.Comment{}, result.Error
	}

	return comment, nil
}
