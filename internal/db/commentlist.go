package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/db/helpers"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetReviewComments(reviewid string, o models.CommentSortOptions) (models.CommentList, error) {
	query := helpers.Format(`SELECT * FROM comments WHERE reviewid = $1`, o)
	rows, err := db.Query(query, reviewid)

	if err != nil {
		log.Printf("%s", err)
		return models.CommentList{}, err
	}

	defer rows.Close()

	return helpers.GetCommentList(rows)
}
