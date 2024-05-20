package db

import (
	"github.com/alex-305/bookbackend/internal/db/helpers"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetReviewComments(username string, reviewid string, o models.CommentSortOptions) ([]models.Comment, error) {
	return db.GetCommentList(username, o, models.NewAP("reviewid", reviewid))
}

func (db *DB) GetCommentList(username string, o models.CommentSortOptions, ap models.AttributeParam) ([]models.Comment, error) {

	var commentList []models.Comment
	q := `SELECT c.commentid, c.reviewid, c.content, c.username,c.post_date, c.likecount, CASE WHEN ulc.username IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked FROM comments c LEFT JOIN user_likes_comment ulc ON c.commentid=ulc.commentid AND ulc.username=$1 WHERE ` + ap.Attribute + `= $2`

	q = helpers.Format(q, o)

	rows, err := db.Query(q, username, ap.Param)

	defer rows.Close()

	if err != nil {
		return []models.Comment{}, err
	}
	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(&comment.ReviewID, &comment.CommentID, &comment.Content, &comment.Username, &comment.Post_date, &comment.LikeCount, &comment.IsLiked)

		if err != nil {
			return []models.Comment{}, err
		}

		commentList = append(commentList, comment)
	}

	return commentList, nil
}
