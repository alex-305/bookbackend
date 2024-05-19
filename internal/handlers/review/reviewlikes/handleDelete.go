package reviewlikes

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewlikesvc"
	"github.com/gorilla/mux"
)

func HandleDelete(w http.ResponseWriter, r *http.Request, d *db.DB) {

	vars := mux.Vars(r)
	reviewid := vars["reviewid"]

	tok, err := helpers.GetToken(r)
	if err != nil {
		http.Error(w, "Could not retrieve token", http.StatusUnauthorized)
		return
	}

	err = reviewlikesvc.Delete(reviewid, tok, d)

	if err != nil {
		http.Error(w, "Could not unlike", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
