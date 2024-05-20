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
	query := `SELECT r.username, r.volumeid, r.reviewid, r.content, r.rating, r.post_date, r.likecount FROM reviews
	CASE WHEN ulr.username IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked FROM reviews r LEFT JOIN user_likes_review ulr ON r.reviewid=ulr.reviewid AND ulr.username=$1 WHERE r.reviewid = $2`

	var rev models.Review
	err := db.QueryRow(query, username, reviewid).Scan(&rev.Username, &rev.VolumeID, &rev.ReviewID, &rev.Content, &rev.Rating, &rev.Post_date, &rev.LikeCount, &rev.IsLiked)

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
