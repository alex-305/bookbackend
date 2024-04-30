package api

import (
	"net/http"

	"github.com/alex-305/bookbackend/handlers/auth"
	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(router *mux.Router) {
	//Auth Routes
	router.HandleFunc("/login", auth.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/signup", auth.HandleCreateAccount).Methods(http.MethodPost)
	//User Routes

}
