package api

import (
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/auth"
	"github.com/alex-305/bookbackend/handlers/review"
	reviewlist "github.com/alex-305/bookbackend/handlers/review/reviewlist"
	"github.com/alex-305/bookbackend/handlers/user"
	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(router *mux.Router) {
	//Auth Routes
	router.HandleFunc("/login", makeHttp(auth.HandleLogin, s.DB)).Methods(http.MethodPost)
	router.HandleFunc("/signup", makeHttp(auth.HandleSignUp, s.DB)).Methods(http.MethodPost)
	//User Routes
	router.HandleFunc("/user/{username}", makeHttp(user.HandleGetUser, s.DB)).Methods(http.MethodGet)
	router.HandleFunc("/user/{username}/description", makeHttp(user.HandlePatchDesc, s.DB)).Methods(http.MethodPatch)
	//Review Routes
	router.HandleFunc("/review", makeHttp(review.HandlePost, s.DB)).Methods(http.MethodPost)
	router.HandleFunc("/review/{reviewid}", makeHttp(review.HandleDelete, s.DB)).Methods(http.MethodDelete)
	router.HandleFunc("/review/{reviewid}", makeHttp(review.HandleGet, s.DB)).Methods(http.MethodGet)
	//Review List Routes
	router.HandleFunc("/user/{username}/reviews", makeHttp(reviewlist.HandleGetUser, s.DB)).Methods(http.MethodGet)
	//router.HandleFunc("/reviews/popular", makeHttp(reviewList.HandleGetPopular, s.DB)).Methods(http.MethodGet)
	//Reply Routes
	// router.HandleFunc("/review/{reviewid}/reply", makeHttp(reply.HandlePost, s.DB)).Methods(http.MethodPost)
	// router.HandleFunc("/review/{reviewid}/reply/")
}

type handlerFunc func(http.ResponseWriter, *http.Request, *db.DB)
type httpFunc func(http.ResponseWriter, *http.Request)

func makeHttp(fn handlerFunc, d *db.DB) httpFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, d)
	})
}
