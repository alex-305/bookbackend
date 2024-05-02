package review

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/helpers"
	"github.com/alex-305/bookbackend/models"
	"github.com/alex-305/bookbackend/services/review"
)

func HandlePostReview(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusForbidden)
		return
	}

	var rev models.Review
	err = json.NewDecoder(r.Body).Decode(&rev)

	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	err = review.PostReview(tok, rev, db)

	if err != nil {
		http.Error(w, "Could not post review", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
