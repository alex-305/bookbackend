package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/authsvc"
)

func HandleLogin(w http.ResponseWriter, r *http.Request, db *db.DB) {
	log.Printf("handleLogin running...")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		http.Error(w, "Could not decode the JSON", http.StatusBadRequest)
		return
	}

	token, _ := helpers.GetToken(r)

	token, err = authsvc.Login(creds, token, db)

	if err != nil && err.Error() == "unable to get credentials" {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}
	tokenJSON := map[string]string{"token": models.String(token)}
	response, err := json.Marshal(tokenJSON)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
