package db

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) PostComment(r models.Comment) (string, error) {
	if r.Content == "" {
		return "", errors.New("content cannot be empty")
	}
	query := `INSERT into comments(username, reviewID, content) VALUES($1, $2, $3) RETURNING commentid`
	var commentid string
	err := db.QueryRow(query, r.Username, r.ReviewID, r.Content).Scan(&commentid)

	if err != nil {
		return "", err
	}
	return commentid, nil
}

func (db *DB) DeleteComment(cid string) error {
	_, err := db.Exec(`DELETE FROM comments WHERE commentid=$1`, cid)

	if err != nil {
		return err
	}
	return nil
}
