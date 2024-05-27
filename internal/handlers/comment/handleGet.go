package comment

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/commentsvc"
	"github.com/gorilla/mux"
)

func HandleGet(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	commentid := models.CommentID(vars["commentid"])

	tok, _ := helpers.GetToken(r)

	comment, err := commentsvc.Get(commentid, tok, d)

	if err != nil {
		http.Error(w, "Could not get comment", http.StatusBadRequest)
		return
	}

	commentJSON, err := json.Marshal(comment)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(commentJSON)

}
