package models

type BookStats struct {
	AvgRating   float32 `json:"avgrating"`
	ReviewCount uint32  `json:"reviewcount"`
	Volumeid    string  `json:"volumeid"`
}
