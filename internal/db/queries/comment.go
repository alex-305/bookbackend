package queries

import (
	"github.com/alex-305/bookbackend/internal/models"
)

func GetComment(ap models.AttributeParam) string {
	return `SELECT c.commentid, c.reviewid, c.content, c.username,c.post_date, c.likecount, CASE WHEN ulc.username IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked FROM comments c LEFT JOIN user_likes_comment ulc ON c.commentid=ulc.commentid AND ulc.username=$1 WHERE c.` + ap.Attribute + `= $2`
}
