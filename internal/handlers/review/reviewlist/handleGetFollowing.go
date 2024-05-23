package reviewlist

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewlistsvc"
)

func HandleGetFollowing(w http.ResponseWriter, r *http.Request, d *db.DB) {
	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusUnauthorized)
		return
	}

	o := getOptions(r)

	reviews, err := reviewlistsvc.GetFollowing(tok, o, d)

	if err != nil {
		http.Error(w, "Could not get following review list", http.StatusBadRequest)
		return
	}

	reviewJSON, err := json.Marshal(reviews)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(reviewJSON)

}
