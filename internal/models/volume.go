package models

type Volume struct {
	ReviewCount uint    `json:"reviewcount"`
	AvgRating   float32 `json:"avgrating"`
	VolumeID    string  `json:"volumeid"`
}

type VolumeID string
