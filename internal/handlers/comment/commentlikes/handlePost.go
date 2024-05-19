package commentlikes

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/services/commentsvc/commentlikesvc"
	"github.com/gorilla/mux"
)

func HandlePost(w http.ResponseWriter, r *http.Request, db *db.DB) {

	vars := mux.Vars(r)
	commentid := vars["commentid"]

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve token", http.StatusUnauthorized)
		return
	}

	err = commentlikesvc.Post(commentid, tok, db)

	if err != nil {
		http.Error(w, "Could not post like", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
