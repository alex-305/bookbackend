package api

import (
	"net/http"

	"github.com/alex-305/bookbackend/handlers/auth"
	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(router *mux.Router) {
	router.HandleFunc("/login", auth.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/user", auth.HandleCreateAccount).Methods(http.MethodPost)
}
