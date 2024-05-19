package reviewlist

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/gorilla/mux"
)

func HandleGetStats(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)
	username := vars["username"]

	stats := db.GetUserReviewStats(username)

	response, err := json.Marshal(stats)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
