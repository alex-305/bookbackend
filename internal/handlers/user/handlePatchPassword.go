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

type pass struct {
	Password string `json:"password"`
}

func HandlePatchPass(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	username := models.Username(vars["username"])

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Token not found", http.StatusUnauthorized)
		return
	}

	var p pass
	err = json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, "Could not decode request body", http.StatusBadRequest)
		return
	}

	err = usersvc.PatchPassword(username, p.Password, tok, d)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not update password", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
