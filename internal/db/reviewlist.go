package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/db/helpers"
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/db/scan"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db DB) GetUserReviewList(userusername string, username string, o models.ReviewSortOptions) ([]models.Review, error) {
	return db.getReviewList(userusername, o, models.NewAP("username", username))
}

func (db DB) GetBookReviewList(username string, volumeID string, o models.ReviewSortOptions) ([]models.Review, error) {
	return db.getReviewList(username, o, models.NewAP("volumeid", volumeID))
}

func (db DB) GetFollowingReviewList(username string, o models.ReviewSortOptions) ([]models.Review, error) {
	q := queries.GetFollowingReviews()
	q += helpers.ListFormat(o)
	rows, err := db.Query(q, username)
	defer rows.Close()
	if err != nil {
		log.Printf("query:%s \n%s", q, err)
		return []models.Review{}, err
	}

	return scan.Reviews(rows)
}

func (db DB) getReviewList(username string, o models.ReviewSortOptions, ap models.AttributeParam) ([]models.Review, error) {

	q := queries.GetReview(ap)
	q += helpers.ListFormat(o)

	rows, err := db.Query(q, username, ap.Param)

	defer rows.Close()

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}

	return scan.Reviews(rows)
}
