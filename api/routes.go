package api

import (
	"net/http"

	"github.com/alex-305/bookbackend/handlers/auth"
	"github.com/alex-305/bookbackend/handlers/review"
	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(router *mux.Router) {
	//Auth Routes
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleLogin(w, r, s.DB)
	}).Methods(http.MethodPost)

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleSignUp(w, r, s.DB)
	}).Methods(http.MethodPost)
	//User Routes
	//Review Routes
	router.HandleFunc("/review", func(w http.ResponseWriter, r *http.Request) {
		review.HandlePostReview(w, r, s.DB)
	}).Methods(http.MethodPost)
	router.HandleFunc("/review", func(w http.ResponseWriter, r *http.Request) {
		review.HandleDeleteReview(w, r, s.DB)
	}).Methods(http.MethodDelete)
}
