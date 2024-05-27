package review

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc"
	"github.com/gorilla/mux"
)

func HandleDelete(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)

	reviewid := models.ReviewID(vars["reviewid"])

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusForbidden)
		return
	}

	err = reviewsvc.Delete(reviewid, tok, db)
	if err != nil {
		http.Error(w, "Could not delete review", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
