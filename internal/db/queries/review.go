package queries

import (
	"github.com/alex-305/bookbackend/internal/models"
	"gorm.io/gorm"
)

func FromReview() string {
	return `reviews r`
}

func ReviewTableName() string {
	return `r`
}

func SelectReview() string {
	return `u.username AS username r.userid, r.volumeid, r.reviewid, r.content, r.rating, r.post_date, r.likecount, CASE WHEN ulr.userid IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked`
}

func JoinReviewLikes() string {
	return ` LEFT JOIN user_likes_review ulr ON r.reviewid=ulr.reviewid AND ulr.userid = ? `
}

func JoinReviewFollowing() string {
	return ` LEFT JOIN user_follows_user ufu ON r.userid=ufu.followedid `
}

func ChangeReviewLikeCount(reviewid models.ReviewID, g *gorm.DB, change int8) *gorm.DB {
	return g.Model(&models.Comment{}).Where(("reviewid = ?"), reviewid).Update("likecount", gorm.Expr("likecount+?", change))
}
