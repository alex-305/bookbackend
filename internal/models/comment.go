package models

import "time"

type Comment struct {
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	CommentID string    `json:"commentid"`
	ReviewID  string    `json:"reviewid"`
	Post_date time.Time `json:"post_date"`
	LikeCount uint      `json:"likecount"`
}

type CommentList struct {
	Comments     []Comment `json:"comments"`
	CommentCount uint      `json:"commentcount"`
}
