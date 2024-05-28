package review

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc"
)

func HandlePost(w http.ResponseWriter, r *http.Request, db *db.DB) {
	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusForbidden)
		return
	}

	var rev models.NewReview
	err = json.NewDecoder(r.Body).Decode(&rev)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	reviewid, err := reviewsvc.Post(rev, tok, db)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not post review", http.StatusBadRequest)
		return
	}

	reviewIDJSON := map[string]models.ReviewID{"reviewid": reviewid}
	response, err := json.Marshal(reviewIDJSON)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
