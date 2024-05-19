package models

import "time"

type Review struct {
	Content   string    `json:"content"`
	Rating    uint8     `json:"rating"`
	Username  string    `json:"username"`
	VolumeID  string    `json:"volumeid"`
	ReviewID  string    `json:"reviewid"`
	Post_date time.Time `json:"post_date"`
	LikeCount uint      `json:"likecount"`
}
type ReviewListStats struct {
	ReviewCount uint `json:"reviewcount"`
	AvgRating   uint `json:"avgrating"`
}

type ReviewStats struct {
	CommentCount uint `json:"commentcount"`
}
