package review

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/reviewsvc"
	"github.com/gorilla/mux"
)

func HandleGet(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	reviewid := models.ReviewID(vars["reviewid"])

	tok, _ := helpers.GetToken(r)

	review, err := reviewsvc.Get(reviewid, tok, d)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not retrieve review", http.StatusNotFound)
		return
	}

	reviewJSON, err := json.Marshal(review)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(reviewJSON)

}
