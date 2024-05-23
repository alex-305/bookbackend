package scan

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func Comments(rows *sql.Rows) ([]models.Comment, error) {
	var commentList []models.Comment
	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(&comment.CommentID, &comment.ReviewID, &comment.Content, &comment.Username, &comment.Post_date, &comment.LikeCount, &comment.IsLiked)

		if err != nil {
			return []models.Comment{}, err
		}

		commentList = append(commentList, comment)
	}
	return commentList, nil
}

func Comment(r *sql.Row) (models.Comment, error) {
	var c models.Comment
	err := r.Scan(&c.CommentID, &c.ReviewID, &c.Content, &c.Username, &c.Post_date, &c.LikeCount, &c.IsLiked)
	return c, err
}
