package review

import (
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/helpers"
)

func HandleDeleteReview(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	_, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not retrieve JWT token", http.StatusForbidden)
		return
	}

}
