package scan

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func User(row *sql.Row) (models.User, error) {
	var user models.User
	err := row.Scan(&user.Username, &user.Email, &user.Description, &user.Join_date, &user.FollowerCount, &user.FollowingCount, &user.IsFollowing)
	return user, err
}
