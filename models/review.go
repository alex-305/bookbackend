package models

import "time"

type Review struct {
	Content   string
	Rating    uint8
	Username  string
	WorksID   string
	ReviewID  string
	Post_date time.Time
}
