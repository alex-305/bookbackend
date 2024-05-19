package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/db/helpers"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db DB) GetUserReviewList(username string, o models.ReviewSortOptions) ([]models.Review, error) {
	query := helpers.Format(`SELECT *, COUNT(*) AS reviewcount FROM reviews WHERE username = $1`, o)

	rows, err := db.Query(query, username)

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}

	defer rows.Close()

	return helpers.GetReviewList(rows)
}

func (db DB) GetUserReviewCount(username string) int {
	query := `SELECT COUNT(*) FROM reviews WHERE username = $1`
	return db.getReviewCount(query, username)
}

func (db DB) GetBookReviewList(volumeID string, o models.ReviewSortOptions) ([]models.Review, error) {
	query := helpers.Format(`SELECT * FROM reviews WHERE volumeid = $1`, o)

	rows, err := db.Query(query, volumeID)

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}

	defer rows.Close()

	return helpers.GetReviewList(rows)
}

func (db DB) GetBookReviewCount(volumeID string) int {
	query := `SELECT COUNT(*) FROM reviews WHERE volumeid = $1`
	return db.getReviewCount(query, volumeID)
}

func (db DB) getReviewCount(query, param string) int {
	var count int
	_ = db.QueryRow(query, param).Scan(&count)
	return count
}
