package helpers

import (
	"database/sql"

	"github.com/alex-305/bookbackend/models"
)

func GetCommentList(rows *sql.Rows) ([]models.Comment, error) {

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(&comment.ReviewID, &comment.CommentID, &comment.Content, &comment.Username, &comment.Post_date, &comment.LikeCount)

		if err != nil {
			return []models.Comment{}, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
