package models

import "time"

type Review struct {
	Content   string    `json:"Content"`
	Rating    uint8     `json:"Rating"`
	Username  string    `json:"Username"`
	WorksID   string    `json:"WorksID"`
	ReviewID  string    `json:"ReviewID"`
	Post_date time.Time `json:"Post_date"`
}
