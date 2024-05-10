package user

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/helpers"
	"github.com/alex-305/bookbackend/services/usersvc"
	"github.com/gorilla/mux"
)

type desc struct {
	Description string `json:"description"`
}

func HandlePatchDesc(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

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

	err = usersvc.PatchDescription(tok, username, d.Description, db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
