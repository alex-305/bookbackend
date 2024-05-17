package db

func (db DB) UpdateUserDescription(username, description string) error {
	query := `UPDATE users SET description = $1 WHERE username = $2`
	_, err := db.Exec(query, description, username)

	if err != nil {
		return err
	}
	return nil
}

func (db DB) UpdateUserPassword(username, password string) error {
	query := `UPDATE users SET password = $1 WHERE username = $2`
	_, err := db.Exec(query, password, username)

	if err != nil {
		return err
	}
	return nil
}
