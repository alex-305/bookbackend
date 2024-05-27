package review

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc"
	"github.com/gorilla/mux"
)

func HandlePut(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	reviewid := models.ReviewID(vars["reviewid"])

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusMethodNotAllowed)
		return
	}
	var rev models.Review
	err = json.NewDecoder(r.Body).Decode(&rev)

	if err != nil {
		http.Error(w, "Could not retrieve body", http.StatusBadRequest)
		return
	}

	err = reviewsvc.Put(reviewid, rev, tok, d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
