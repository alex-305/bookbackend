package db

import (
	"log"

	"github.com/alex-305/bookbackend/models"
)

func (db *DB) PostComment(r models.Comment) (string, error) {
	query := `INSERT into comments(username, reviewID, content) VALUES($1, $2, $3) RETURNING commentid`
	var commentid string
	err := db.QueryRow(query, r.Username, r.ReviewID, r.Content).Scan(&commentid)

	log.Printf("%+v", r)

	if err != nil {
		return "", err
	}
	return commentid, nil
}
