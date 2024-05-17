package commentlist

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/services/commentsvc/commentlistsvc"
	"github.com/gorilla/mux"
)

func HandleGetReview(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	reviewid := vars["reviewid"]

	options := getOptions(r)

	comments, err := commentlistsvc.GetReview(reviewid, options, d)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not retrieve comments", http.StatusNotFound)
		return
	}

	commentsJSON, err := json.Marshal(comments)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(commentsJSON)
}
