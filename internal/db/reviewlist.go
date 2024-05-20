package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/db/helpers"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db DB) GetUserReviewList(userusername string, username string, o models.ReviewSortOptions) ([]models.Review, error) {
	return db.GetReviewList(userusername, o, models.NewAP("username", username))
}

func (db DB) GetBookReviewList(username string, volumeID string, o models.ReviewSortOptions) ([]models.Review, error) {
	return db.GetReviewList(username, o, models.NewAP("volumeid", volumeID))
}

func (db DB) GetReviewList(username string, o models.ReviewSortOptions, ap models.AttributeParam) ([]models.Review, error) {

	var reviewList []models.Review
	q := `SELECT r.username, r.volumeid, r.reviewid, r.content, r.rating, r.post_date, r.likecount, CASE WHEN ulr.username IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked FROM reviews r LEFT JOIN user_likes_review ulr ON r.reviewid=ulr.reviewid AND ulr.username=$1 WHERE r.` + ap.Attribute + `= $2`

	q = helpers.Format(q, o)

	rows, err := db.Query(q, username, ap.Param)

	defer rows.Close()

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}

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
