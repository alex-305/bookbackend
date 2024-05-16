package book

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/services/booksvc"
	"github.com/gorilla/mux"
)

func HandleGet(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	volumeid := vars["volumeid"]

	bs, err := booksvc.Get(volumeid, d)

	if err != nil {
		http.Error(w, "Could not retrieve stats", http.StatusNotFound)
		return
	}

	bsJSON, err := json.Marshal(bs)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bsJSON)
}
