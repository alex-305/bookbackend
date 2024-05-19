package helpers

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func GetCommentList(rows *sql.Rows) ([]models.Comment, error) {

	var commentList []models.Comment

	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(&comment.ReviewID, &comment.CommentID, &comment.Content, &comment.Username, &comment.Post_date, &comment.LikeCount)

		if err != nil {
			return []models.Comment{}, err
		}

		commentList = append(commentList, comment)
	}

	return commentList, nil
}
