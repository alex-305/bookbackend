package reviewstats

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc/reviewstatssvc"
	"github.com/gorilla/mux"
)

func HandleGet(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	reviewid := models.ReviewID(vars["reviewid"])

	stats, err := reviewstatssvc.GetReviewStats(reviewid, d)

	if err != nil {
		http.Error(w, "Could not get review stats", http.StatusBadRequest)
		return
	}

	statsJSON, err := json.Marshal(stats)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(statsJSON)

}
