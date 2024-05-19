package reviewstats

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewstatssvc"
	"github.com/gorilla/mux"
)

func HandleGet(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	reviewid := vars["reviewid"]

	stats := reviewstatssvc.Get(reviewid, d)

	statsJSON, err := json.Marshal(stats)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(statsJSON)

}
