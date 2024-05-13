package reviewlist

import (
	"net/http"
	"strconv"

	"github.com/alex-305/bookbackend/models"
)

func getOptions(r *http.Request) models.SortOptions {
	q := r.URL.Query()

	o := models.SortOptions{
		By:        models.PostDate,
		Direction: models.Descending,
		Limit:     50,
		Page:      0,
	}
	limitStr := q.Get("amount")
	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err == nil {
		o.Limit = uint(limit)
	}
	//limit, err = strconv.ParseUint(limit, 10, 32)
	pageStr := q.Get("page")
	page, err := strconv.ParseUint(pageStr, 10, 32)

	if err == nil {
		o.Page = uint(page)
	}

	o.By = models.GetSortBy(q.Get("sort_by"))

	o.Direction = models.GetSortDirection(q.Get("sort_dir"))

	return o
}
