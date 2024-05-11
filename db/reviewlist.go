package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alex-305/bookbackend/models"
)

func (db DB) GetUserReviewList(username string, o models.SortOptions) ([]models.Review, error) {
	query := fmt.Sprintf(`SELECT * FROM reviews WHERE username = $1 ORDER BY %s %s LIMIT %d OFFSET %d;`, o.By, o.Direction, o.Limit, o.Page*o.Limit)

	log.Printf("%s", query)

	rows, err := db.Query(query, username)

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}
	return db.GetReviewList(rows)
}

func (db DB) GetUserReviewCount(username string) int {
	query := `SELECT COUNT(*) FROM reviews WHERE username = $1`
	var count int
	_ = db.QueryRow(query, username).Scan(&count)

	return count
}

func (db DB) GetReviewList(rows *sql.Rows) ([]models.Review, error) {

	defer rows.Close()

	var reviews []models.Review

	for rows.Next() {
		var review models.Review

		var rating *uint8
		var content sql.NullString

		err := rows.Scan(&review.Username, &review.WorksID, &review.ReviewID, &content, &rating, &review.Post_date)

		if err != nil {
			return []models.Review{}, err
		}

		if rating != nil {
			review.Rating = rating
		}

		if content.Valid {
			review.Content = content.String
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}
