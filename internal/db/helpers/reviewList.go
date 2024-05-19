package helpers

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func GetReviewList(rows *sql.Rows) ([]models.Review, error) {

	var reviews []models.Review

	for rows.Next() {
		var review models.Review

		var content sql.NullString

		err := rows.Scan(&review.Username, &review.VolumeID, &review.ReviewID, &content, &review.Rating, &review.Post_date, &review.LikeCount)

		if err != nil {
			return []models.Review{}, err
		}

		if content.Valid {
			review.Content = content.String
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
