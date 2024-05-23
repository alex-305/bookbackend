package models

type SortOptions interface {
	GetBy() string
	GetDirection() string
	GetLimit() uint
	GetPage() uint
}
type ReviewSortOptions struct {
	By        ReviewSortBy
	Direction SortDirection
	Limit     uint
	Page      uint
}

type CommentSortOptions struct {
	By        CommentSortBy
	Direction SortDirection
	Limit     uint
	Page      uint
}

type FollowOptions struct {
	Limit uint
	Page  uint
}

type ReviewSortBy string

const (
	RevPostDate ReviewSortBy = "post_date"
	RevLikes    ReviewSortBy = "likecount"
	RevRating   ReviewSortBy = "rating"
)

type CommentSortBy string

const (
	ComLikes    CommentSortBy = "likecount"
	ComPostDate CommentSortBy = "post_date"
)

type SortDirection string

const (
	Descending SortDirection = "DESC"
	Ascending  SortDirection = "ASC"
)

func ParseSortDirection(s string) SortDirection {
	dir := Descending
	if s == "ascending" {
		dir = Ascending
	}
	return dir
}

func ParseReviewSortBy(s string) ReviewSortBy {
	switch s {
	case "rating":
		return RevRating
	case "likes":
		return RevLikes
	default:
		return RevPostDate
	}
}

func ParseCommentSortBy(s string) CommentSortBy {
	switch s {
	case "likes":
		return ComLikes
	default:
		return ComPostDate
	}
}

func (o ReviewSortOptions) GetBy() string {
	return string(o.By)
}
func (o CommentSortOptions) GetBy() string {
	return string(o.By)
}
func (o ReviewSortOptions) GetDirection() string {
	return string(o.Direction)
}
func (o CommentSortOptions) GetDirection() string {
	return string(o.Direction)
}
func (o ReviewSortOptions) GetLimit() uint {
	return uint(o.Limit)
}
func (o CommentSortOptions) GetLimit() uint {
	return uint(o.Limit)
}
func (o ReviewSortOptions) GetPage() uint {
	return uint(o.Page)
}
func (o CommentSortOptions) GetPage() uint {
	return uint(o.Page)
}
