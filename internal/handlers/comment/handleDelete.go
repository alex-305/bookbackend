package comment

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/commentsvc"
	"github.com/gorilla/mux"
)

func HandleDelete(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	commentid := models.CommentID(vars["commentid"])

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve token", http.StatusUnauthorized)
		return
	}

	err = commentsvc.Delete(commentid, tok, d)

	if err != nil {
		http.Error(w, "Could not post comment", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
