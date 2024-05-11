package reviewlist

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
	"github.com/alex-305/bookbackend/services/reviewsvc/reviewlistsvc"
	"github.com/gorilla/mux"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	options := getOptions(r)

	reviews, err := reviewlistsvc.GetUser(username, options, db)

	if err != nil {
		http.Error(w, "Could not get user list", http.StatusBadRequest)
		return
	}

	reviewsJSON, err := json.Marshal(reviews)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(reviewsJSON)
}

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
