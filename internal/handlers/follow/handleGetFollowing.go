package follow

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/services/followsvc/followlistsvc"
	"github.com/gorilla/mux"
)

func HandleGetFollowing(w http.ResponseWriter, r *http.Request, db *db.DB) {
	vars := mux.Vars(r)
	username := models.Username(vars["username"])

	tok, _ := helpers.GetToken(r)

	o := getOptions(r)

	followerList, err := followlistsvc.Following(username, tok, o, db)

	if err != nil {
		http.Error(w, "Could not get following list", http.StatusUnauthorized)
		return
	}

	followerJSON, err := json.Marshal(followerList)

	if err != nil {
		http.Error(w, "Could not marshal json for response", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(followerJSON)
}
