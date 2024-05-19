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

type ReviewList struct {
	Reviews     []Review `json:"reviews"`
	ReviewCount uint     `json:"reviewcount"`
}
