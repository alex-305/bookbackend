package db

import (
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetReviewComments(userID models.UserID, reviewid models.ReviewID, o models.CommentSortOptions) ([]models.Comment, error) {
	return db.getCommentList(userID, o, models.NewAP("reviewid", string(reviewid)))
}

func (db *DB) getCommentList(userID models.UserID, o models.CommentSortOptions, ap models.AttributeParam) ([]models.Comment, error) {
	var comments []models.Comment
	result :=
		db.Table(queries.FromComment()).
			Select(queries.SelectComment()).
			Joins(queries.JoinCommentLikes(), userID).
			//Joins(queries.JoinForUsername(queries.CommentTableName())).
			Where(ap.Attribute, ap.Param).
			Order(string(o.By) + " " + string(o.Direction)).
			Limit(int(o.Limit)).
			Offset(int(o.Limit * o.Page)).
			Scan(&comments)

	if result.Error != nil {
		return []models.Comment{}, result.Error
	}

	return comments, nil
}
