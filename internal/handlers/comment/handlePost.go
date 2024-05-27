package comment

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/commentsvc"
	"github.com/gorilla/mux"
)

func HandlePost(w http.ResponseWriter, r *http.Request, d *db.DB) {

	var c models.Comment

	vars := mux.Vars(r)
	c.ReviewID = models.ReviewID(vars["reviewid"])

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Token not found", http.StatusUnauthorized)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, "Could not decode json", http.StatusBadRequest)
		return
	}

	commentid, err := commentsvc.Post(c, tok, d)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not post comment", http.StatusInternalServerError)
		return
	}

	commentidJSON := map[string]models.CommentID{"commentid": commentid}
	response, err := json.Marshal(commentidJSON)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
