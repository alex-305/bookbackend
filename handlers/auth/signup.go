package auth

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/auth"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func HandleSignUp(w http.ResponseWriter, r *http.Request, db *db.DB) {
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

	token, err := auth.SignUp(creds, db)
	if err != nil {
		http.Error(w, "Could not create account", http.StatusBadRequest)
		return
	}
	tokenJSON := map[string]string{"token": token}
	response, err := json.Marshal(tokenJSON)

	if err != nil {
		http.Error(w, "Could not parse json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
