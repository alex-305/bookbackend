package db

import (
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) PostReview(username string, rev models.Review) (string, error) {
	query := `INSERT INTO reviews(username, volumeid, content, rating) VALUES($1, $2, $3, $4) RETURNING reviewID;`

	var reviewid string
	err := db.QueryRow(query, username, rev.VolumeID, rev.Content, rev.Rating).Scan(&reviewid)

	if err != nil {
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

func (db *DB) GetReview(username, reviewid string) (models.Review, error) {
	query := `SELECT username, volumeid, reviewid, content, rating, post_date, likecount FROM reviews WHERE reviewid = $1`

	var rev models.Review
	err := db.QueryRow(query, reviewid).Scan(&rev.Username, &rev.VolumeID, &rev.ReviewID, &rev.Content, &rev.Rating, &rev.Post_date, &rev.LikeCount, &rev.IsLiked)

	if err != nil {
		return models.Review{}, err
	}

	return rev, nil
}

func (db *DB) UpdateReview(reviewid string, r models.Review) error {
	query := `UPDATE reviews SET content = $1, rating = $2 WHERE reviewid = $3`
	_, err := db.Exec(query, r.Content, r.Rating, reviewid)

	if err != nil {
		return err
	}
	return nil
}
