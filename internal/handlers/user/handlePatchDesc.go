package user

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/usersvc"
	"github.com/gorilla/mux"
)

type desc struct {
	Description string `json:"description"`
}

func HandlePatchDesc(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)
	username := models.Username(vars["username"])

	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusMethodNotAllowed)
		return
	}
	var d desc
	err = json.NewDecoder(r.Body).Decode(&d)

	if err != nil {
		http.Error(w, "Could not retrieve body", http.StatusBadRequest)
		return
	}

	err = usersvc.PatchDescription(username, d.Description, tok, db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
