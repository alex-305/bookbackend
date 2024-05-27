package reviewlikes

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewlikesvc"
	"github.com/gorilla/mux"
)

func HandlePost(w http.ResponseWriter, r *http.Request, d *db.DB) {

	vars := mux.Vars(r)
	reviewid := models.ReviewID(vars["reviewid"])

	tok, err := helpers.GetToken(r)
	if err != nil {
		http.Error(w, "Could not retrieve token", http.StatusUnauthorized)
		return
	}

	err = reviewlikesvc.Post(reviewid, tok, d)

	if err != nil {
		http.Error(w, "Could not like", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
