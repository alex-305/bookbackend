package db

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetUser(userusername, username string) (models.User, error) {
	row := db.QueryRow("SELECT u.username, u.email, u.description, u.join_date, u.followercount, u.followingcount, CASE WHEN ufu.follower IS NOT NULL THEN TRUE ELSE FALSE END AS isfollowing FROM users u LEFT JOIN user_follows_user ufu ON u.username=ufu.followed AND ufu.follower=$1 WHERE username = $2", userusername, username)

	var user models.User
	err := row.Scan(&user.Username, &user.Email, &user.Description, &user.Join_date, &user.FollowerCount, &user.IsFollowing)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (db *DB) GetCredentials(username string) (models.Credentials, error) {
	row := db.QueryRow("SELECT username, email, password FROM users WHERE username = $1", username)
	var creds models.Credentials

	err := row.Scan(&creds.Username, &creds.Email, &creds.Password)

	if err != nil {
		return models.Credentials{}, errors.New("unable to get credentials")
	}

	return creds, nil
}

func (db *DB) CreateUser(creds models.Credentials) error {
	query := `INSERT INTO users(username, password, email) VALUES($1, $2, $3);`
	_, err := db.Exec(query, creds.Username, creds.Password, creds.Email)

	if err != nil {
		return err
	}
	return nil
}
