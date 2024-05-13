package review

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/helpers"
	"github.com/alex-305/bookbackend/models"
	"github.com/alex-305/bookbackend/services/reviewsvc"
)

func HandlePost(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusForbidden)
		return
	}

	var rev models.Review
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

	reviewIDJSON := map[string]string{"reviewid": reviewid}
	response, err := json.Marshal(reviewIDJSON)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
