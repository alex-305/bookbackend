package reviewlist

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewlistsvc/reviewliststatssvc"
	"github.com/gorilla/mux"
)

func HandleGetStats(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)
	username := models.Username(vars["username"])

	stats, err := reviewliststatssvc.GetUserStats(username, db)

	if err != nil {
		http.Error(w, "Could not get user review stats", http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(stats)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
