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

// func (db DB) GetFollowingReviewList(userusername, username string) ([]models.Review, error) {
// 	q := `SELECT uf1.follower FROM user_follows_user AS uf1 lEFT JOIN user_follows_user AS uf2 ON uf1.follower=uf2.followed AND uf2.follower=$1 WHERE uf1.followed=$2 ORDER BY uf2.follower IS NULL, uf1.follower`
// 	q2 := `SELECT * FROM user_follows_user WHERE followed='ted' ORDER BY followdate`
// 	return []models.Review{}, nil
// }

func (db DB) getReviewList(username string, o models.ReviewSortOptions, ap models.AttributeParam) ([]models.Review, error) {

	q := queries.GetReview(ap)
	q = helpers.ListFormat(q, o)

	rows, err := db.Query(q, username, ap.Param)

	defer rows.Close()

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}

	return scan.Reviews(rows)
}
