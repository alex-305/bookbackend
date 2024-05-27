package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/usersvc"
	"github.com/gorilla/mux"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)
	username := models.Username(vars["username"])
	tok, _ := helpers.GetToken(r)
	user, err := usersvc.Get(username, tok, db)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}
