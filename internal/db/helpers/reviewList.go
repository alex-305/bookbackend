package helpers

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func GetReviewList(rows *sql.Rows) (models.ReviewList, error) {

	var reviewList models.ReviewList
	count := uint(0)
	for rows.Next() {
		count++
		var review models.Review

		var content sql.NullString

		err := rows.Scan(&review.Username, &review.VolumeID, &review.ReviewID, &content, &review.Rating, &review.Post_date, &review.LikeCount)

		if err != nil {
			return models.ReviewList{}, err
		}

		if content.Valid {
			review.Content = content.String
		}

		reviewList.Reviews = append(reviewList.Reviews, review)
	}
	reviewList.ReviewCount = count

	return reviewList, nil
}
