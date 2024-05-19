package helpers

import (
	"database/sql"

	"github.com/alex-305/bookbackend/internal/models"
)

func GetCommentList(rows *sql.Rows) (models.CommentList, error) {

	var commentList models.CommentList
	count := uint(0)

	for rows.Next() {
		count++
		var comment models.Comment

		err := rows.Scan(&comment.ReviewID, &comment.CommentID, &comment.Content, &comment.Username, &comment.Post_date, &comment.LikeCount)

		if err != nil {
			return models.CommentList{}, err
		}

		commentList.Comments = append(commentList.Comments, comment)
	}
	commentList.CommentCount = count

	return commentList, nil
}
