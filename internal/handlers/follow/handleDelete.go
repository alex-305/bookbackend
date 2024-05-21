package follow

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/helpers"
	"github.com/alex-305/bookbackend/internal/services/followsvc"
	"github.com/gorilla/mux"
)

func HandleDelete(w http.ResponseWriter, r *http.Request, d *db.DB) {
	vars := mux.Vars(r)
	username := vars["username"]
	tok, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve token", http.StatusUnauthorized)
		return
	}

	err = followsvc.DeleteFollow(username, tok, d)

	if err != nil {
		http.Error(w, "Could not post follow", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}
