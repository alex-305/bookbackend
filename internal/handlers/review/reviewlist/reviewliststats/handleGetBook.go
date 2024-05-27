package reviewliststats

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewlistsvc/reviewliststatssvc"
	"github.com/gorilla/mux"
)

func HandleGetBook(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	volumeid := models.VolumeID(vars["volumeid"])

	bs, err := reviewliststatssvc.GetBook(volumeid, d)

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
