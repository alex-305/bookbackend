package commentlikes

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/services/commentsvc/commentlikesvc"
	"github.com/gorilla/mux"
)

func HandleDelete(w http.ResponseWriter, r *http.Request, db *db.DB) {

	vars := mux.Vars(r)
	commentid := vars["commentid"]

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve token", http.StatusUnauthorized)
		return
	}

	err = commentlikesvc.Delete(commentid, tok, db)

	if err != nil {
		http.Error(w, "Could not unlike", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
