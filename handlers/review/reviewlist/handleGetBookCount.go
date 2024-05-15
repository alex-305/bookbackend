package reviewlist

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/gorilla/mux"
)

func HandleGetBookCount(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)
	volumeid := vars["volumeid"]

	count := db.GetBookReviewCount(volumeid)

	countJSON := map[string]int{"count": count}
	response, err := json.Marshal(countJSON)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
