package queries

import (
	"github.com/alex-305/bookbackend/internal/models"
	"gorm.io/gorm"
)

func FromComment() string {
	return `comments c`
}

func CommentTableName() string {
	return `c`
}

func SelectComment() string {
	return `c.commentid, c.reviewid, c.content, u.username, c.post_date, c.likecount, CASE WHEN ulc.username IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked`
}

func JoinCommentLikes() string {
	return ` LEFT JOIN user_likes_comment ulc ON c.commentid = ulc.commentid AND ulc.userid = ? `
}

func ChangeCommentLikeCount(commentid models.CommentID, g *gorm.DB, change int8) *gorm.DB {
	return g.Model(&models.Comment{}).Where(("commentid = ?"), commentid).Update("likecount", gorm.Expr("likecount+?", change))
}
