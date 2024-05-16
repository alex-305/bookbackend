package db

import "github.com/alex-305/bookbackend/models"

func (db DB) GetBookStats(volumeid string) (models.BookStats, error) {
	row := db.QueryRow(`SELECT AVG(rating) AS avgrating, COUNT(*) as reviewcount FROM reviews WHERE volumeid = $1`, volumeid)

	var bs models.BookStats
	bs.Volumeid = volumeid
	err := row.Scan(&bs.AvgRating, &bs.ReviewCount)
	if err != nil {
		return models.BookStats{}, err
	}
	return bs, nil
}
