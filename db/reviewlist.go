package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alex-305/bookbackend/models"
)

func (db DB) GetUserReviewList(username string, o models.SortOptions) ([]models.Review, error) {
	query := fmt.Sprintf(`SELECT * FROM reviews WHERE username = $1 ORDER BY %s %s LIMIT $2 OFFSET $3;`, o.By, o.Direction)

	rows, err := db.Query(query, username, o.Limit, o.Page*o.Limit)

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}
	return db.GetReviewList(rows)
}

func (db DB) GetReviewList(rows *sql.Rows) ([]models.Review, error) {

	defer rows.Close()

	var reviews []models.Review

	for rows.Next() {
		var review models.Review

		err := rows.Scan(&review.Username, &review.WorksID, &review.ReviewID, &review.Content, &review.Rating, &review.Post_date)

		if err != nil {
			return []models.Review{}, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
