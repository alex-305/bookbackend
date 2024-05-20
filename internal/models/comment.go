package models

import "time"

type Comment struct {
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	CommentID string    `json:"commentid"`
	ReviewID  string    `json:"reviewid"`
	Post_date time.Time `json:"post_date"`
	LikeCount uint      `json:"likecount"`
	IsLiked   bool      `json:"isliked"`
}

type CommentStats struct {
	CommentCount uint `json:"commentcount"`
}
