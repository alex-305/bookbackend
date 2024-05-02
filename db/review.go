package db

import "github.com/alex-305/bookbackend/models"

func (db *DB) PostReview(username string, rev models.Review) error {
	query := `INSERT INTO reviews(username, worksID, content, rating) VALUES($1, $2, $3);`
	_, err := db.Exec(query, username, rev.WorksID, rev.Content, rev.Rating)

	if err != nil {
		return err
	}

	return nil
}
