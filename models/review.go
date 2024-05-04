package models

import "time"

type Review struct {
	Content   string    `json:"content"`
	Rating    uint8     `json:"rating"`
	Username  string    `json:"username"`
	WorksID   string    `json:"worksid"`
	ReviewID  string    `json:"reviewid"`
	Post_date time.Time `json:"post_date"`
}
