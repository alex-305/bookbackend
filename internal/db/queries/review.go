package queries

import "github.com/alex-305/bookbackend/internal/models"

func GetReview(ap models.AttributeParam) string {
	return `SELECT r.username, r.volumeid, r.reviewid, r.content, r.rating, r.post_date, r.likecount,
	CASE WHEN ulr.username IS NOT NULL THEN TRUE ELSE FALSE END AS isLiked FROM reviews r LEFT JOIN user_likes_review ulr ON r.reviewid=ulr.reviewid AND ulr.username=$1 WHERE r.` + ap.Attribute + `= $2`
}

func GetFollowingReviews() string {
	return `SELECT reviewid, username FROM reviews r LEFT JOIN user_follows_user ufu ON r.username=ufu.followed WHERE ufu.follower=$1`
}
