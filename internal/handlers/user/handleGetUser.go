package user

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/services/usersvc"
	"github.com/gorilla/mux"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	user, err := usersvc.Get(username, db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}
