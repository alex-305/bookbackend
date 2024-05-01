package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/auth"
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/helpers"
	"github.com/alex-305/bookbackend/models"
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

	token, err = auth.Login(creds, token, db)
	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not log user in", http.StatusBadRequest)
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
