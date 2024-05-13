package review

import (
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/helpers"
	"github.com/alex-305/bookbackend/services/reviewsvc"
	"github.com/gorilla/mux"
)

func HandleDelete(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)

	reviewid := vars["reviewid"]

	err = reviewsvc.Delete(reviewid, tok, db)
	if err != nil {
		http.Error(w, "Could not delete review", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
