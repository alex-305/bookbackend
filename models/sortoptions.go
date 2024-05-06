package models

type SortOptions struct {
	By        SortBy
	Direction SortDirection
	Limit     uint
	Page      uint
}

type SortBy string

const (
	PostDate SortBy = "post_date"
	Rating   SortBy = "rating"
)

type SortDirection string

const (
	Descending SortDirection = "DESC"
	Ascending  SortDirection = "ASC"
)

func GetSortDirection(s string) SortDirection {
	dir := Descending
	if s == "ascending" {
		dir = Ascending
	}
	return dir
}

func GetSortBy(s string) SortBy {
	by := PostDate

	if s == "rating" {
		by = Rating
	}

	return by
}
