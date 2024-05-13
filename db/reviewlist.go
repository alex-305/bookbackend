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
		rows.Close()
		return []models.Review{}, err
	}
	return db.getReviewList(rows)
}

func (db DB) GetUserReviewCount(username string) int {
	query := `SELECT COUNT(*) FROM reviews WHERE username = $1`
	return db.getReviewCount(query, username)
}

func (db DB) GetBookReviewList(worksID string, o models.SortOptions) ([]models.Review, error) {
	query := fmt.Sprintf(`SELECT * FROM reviews WHERE worksID = $1 ORDER BY %s %s LIMIT %d OFFSET %d`, o.By, o.Direction, o.Limit, o.Page*o.Limit)

	rows, err := db.Query(query, worksID)
	if err != nil {
		rows.Close()
		return []models.Review{}, err
	}
	return db.getReviewList(rows)
}

func (db DB) GetBookReviewCount(worksID string) int {
	query := `SELECT COUNT(*) FROM reviews WHERE worksID = $1`
	return db.getReviewCount(query, worksID)
}

func (db DB) getReviewCount(query, param string) int {
	var count int
	_ = db.QueryRow(query, param).Scan(&count)
	return count
}

func (db DB) getReviewList(rows *sql.Rows) ([]models.Review, error) {

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
