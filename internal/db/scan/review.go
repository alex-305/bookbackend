package scan

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func Reviews(rows *sql.Rows) ([]models.Review, error) {
	var reviewList []models.Review
	for rows.Next() {
		var review models.Review

		err := rows.Scan(&review.Username, &review.VolumeID, &review.ReviewID, &review.Content, &review.Rating, &review.Post_date, &review.LikeCount, &review.IsLiked)

		if err != nil {
			return []models.Review{}, err
		}

		reviewList = append(reviewList, review)
	}

	return reviewList, nil
}

func Review(row *sql.Row) (models.Review, error) {
	var review models.Review

	err := row.Scan(&review.Username, &review.VolumeID, &review.ReviewID, &review.Content, &review.Rating, &review.Post_date, &review.LikeCount, &review.IsLiked)

	return review, err
}
