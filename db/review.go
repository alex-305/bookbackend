package db

import (
	"github.com/alex-305/bookbackend/models"
)

func (db *DB) PostReview(username string, rev models.Review) (string, error) {
	tx, err := db.Begin()

	if err != nil {
		return "", err
	}

	query := `INSERT INTO reviews(username, worksID, content, rating) VALUES($1, $2, $3) RETURNING reviewID;`

	var reviewid string
	err = tx.QueryRow(query, username, rev.WorksID, rev.Content, rev.Rating).Scan(&reviewid)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return reviewid, nil
}

func (db *DB) DeleteReview(reviewid string) error {
	query := `DELETE FROM reviews WHERE reviewID=$1`
	_, err := db.Exec(query, reviewid)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetReview(reviewid string) (models.Review, error) {
	query := `SELECT username, worksid, reviewid, content, rating, post_date FROM reviews WHERE reviewid = $1`

	var rev models.Review
	err := db.QueryRow(query, reviewid).Scan(&rev)

	if err != nil {
		return models.Review{}, err
	}

	return rev, nil
}
