package db

import (
	"log"

	"github.com/alex-305/bookbackend/db/helpers"
	"github.com/alex-305/bookbackend/models"
)

func (db *DB) GetReviewComments(reviewid string, o models.CommentSortOptions) ([]models.Comment, error) {
	query := helpers.Format(`SELECT * FROM comments WHERE reviewid = $1`, o)
	rows, err := db.Query(query, reviewid)

	if err != nil {
		log.Printf("%s", err)
		return []models.Comment{}, err
	}

	defer rows.Close()

	return helpers.GetCommentList(rows)
}
