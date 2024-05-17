package reviewlist

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewlistsvc"
	"github.com/gorilla/mux"
)

func HandleGetBook(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	volumeid := vars["volumeid"]

	options := getOptions(r)

	reviews, err := reviewlistsvc.GetBook(volumeid, options, d)

	if err != nil {
		http.Error(w, "Could not get review list", http.StatusBadRequest)
		return
	}

	reviewsJSON, err := json.Marshal(reviews)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(reviewsJSON)
}
