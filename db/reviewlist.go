package db

import (
	"database/sql"

	"github.com/alex-305/bookbackend/models"
)

func (db DB) GetUserReviewList(username string, o models.SortOptions) ([]models.Review, error) {
	query := `SELECT * FROM reviews WHERE username = $1 ORDER BY $2 $3 LIMIT $4 OFFSET $5;`

	rows, err := db.Query(query, username, o.By, o.Direction, o.Limit, o.Page*o.Limit)

	if err != nil {
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
