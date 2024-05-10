package db

func (db DB) UpdateUserDescription(username, description string) error {
	query := `UPDATE FROM users SET description = $1 WHERE username $2`
	_, err := db.Exec(query, description, username)

	if err != nil {
		return err
	}
	return nil
}
