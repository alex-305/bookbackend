package db

import (
	"github.com/alex-305/bookbackend/internal/db/helpers"
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/db/scan"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetReviewComments(username string, reviewid string, o models.CommentSortOptions) ([]models.Comment, error) {
	return db.getCommentList(username, o, models.NewAP("reviewid", reviewid))
}

func (db *DB) getCommentList(username string, o models.CommentSortOptions, ap models.AttributeParam) ([]models.Comment, error) {
	q := queries.GetComment(ap)

	q += helpers.ListFormat(o)

	rows, err := db.Query(q, username, ap.Param)

	defer rows.Close()

	if err != nil {
		return []models.Comment{}, err
	}

	return scan.Comments(rows)
}
