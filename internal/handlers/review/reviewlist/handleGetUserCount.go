package reviewlist

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/gorilla/mux"
)

func HandleGetUserCount(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)
	username := vars["username"]

	count := db.GetUserReviewCount(username)

	countJSON := map[string]int{"count": count}
	response, err := json.Marshal(countJSON)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
