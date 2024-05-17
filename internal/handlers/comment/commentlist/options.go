package commentlist

import (
	"net/http"
	"strconv"

	"github.com/alex-305/bookbackend/internal/models"
)

func getOptions(r *http.Request) models.CommentSortOptions {
	q := r.URL.Query()

	o := models.CommentSortOptions{
		By:        models.ComPostDate,
		Direction: models.Descending,
		Limit:     50,
		Page:      0,
	}
	limitStr := q.Get("amount")
	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err == nil {
		o.Limit = uint(limit)
	}

	pageStr := q.Get("page")
	page, err := strconv.ParseUint(pageStr, 10, 32)

	if err == nil {
		o.Page = uint(page)
	}

	o.By = models.ParseCommentSortBy(q.Get("sort_by"))

	o.Direction = models.ParseSortDirection(q.Get("sort_dir"))

	return o
}
