package api

import (
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/auth"
	"github.com/alex-305/bookbackend/handlers/review"
	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(router *mux.Router) {
	//Auth Routes
	router.HandleFunc("/login", makeHttp(auth.HandleLogin, s.DB)).Methods(http.MethodPost)
	router.HandleFunc("/signup", makeHttp(auth.HandleSignUp, s.DB)).Methods(http.MethodPost)
	//User Routes
	//Review Routes
	router.HandleFunc("/review", makeHttp(review.HandlePost, s.DB)).Methods(http.MethodPost)
	router.HandleFunc("/review/{reviewid}", makeHttp(review.HandleDelete, s.DB)).Methods(http.MethodDelete)
	router.HandleFunc("/review/{reviewid}", makeHttp(review.HandleGet, s.DB)).Methods(http.MethodGet)
}

type handlerFunc func(http.ResponseWriter, *http.Request, *db.DB)
type httpFunc func(http.ResponseWriter, *http.Request)

func makeHttp(fn handlerFunc, d *db.DB) httpFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, d)
	})
}
