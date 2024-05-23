package scan

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func Follow(row *sql.Row) (models.UserFollow, error) {
	var f models.UserFollow
	err := row.Scan(&f.Username, &f.IsFollowing)
	return f, err
}

func Follows(rows *sql.Rows) ([]models.UserFollow, error) {
	var followList []models.UserFollow
	for rows.Next() {
		var follow models.UserFollow

		err := rows.Scan(&follow.Username, &follow.IsFollowing)

		if err != nil {
			return []models.UserFollow{}, err
		}

		followList = append(followList, follow)
	}
	return followList, nil
}
