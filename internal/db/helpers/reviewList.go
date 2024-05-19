package helpers

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func GetReviewList(rows *sql.Rows) ([]models.Review, error) {

	var reviewList []models.Review
	count := uint(0)
	for rows.Next() {
		count++
		var review models.Review

		var content sql.NullString

		err := rows.Scan(&review.Username, &review.VolumeID, &review.ReviewID, &content, &review.Rating, &review.Post_date, &review.LikeCount)

		if err != nil {
			return []models.Review{}, err
		}

		if content.Valid {
			review.Content = content.String
		}

		reviewList = append(reviewList, review)
	}

	return reviewList, nil
}
