package reviewlist

import (
	"encoding/json"

	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewlistsvc"
	"github.com/gorilla/mux"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	username := models.Username(vars["username"])

	options := getOptions(r)

	tok, _ := helpers.GetToken(r)

	reviews, err := reviewlistsvc.GetUser(username, tok, options, db)

	if err != nil {
		http.Error(w, "Could not get review list", http.StatusBadRequest)
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
