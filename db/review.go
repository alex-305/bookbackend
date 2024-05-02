package db

func (db *DB) PostReview(username, content string, rating uint8) error {
	query := `INSERT INTO reviews(username,content, rating) VALUES($1, $2, $3);`
	_, err := db.Exec(query, username, content, rating)

	if err != nil {
		return err
	}

	return nil
}
