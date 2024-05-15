package api

import (
	"net/http"

	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/handlers/auth"
	"github.com/alex-305/bookbackend/handlers/comment"
	"github.com/alex-305/bookbackend/handlers/comment/commentlist"
	"github.com/alex-305/bookbackend/handlers/review"
	"github.com/alex-305/bookbackend/handlers/review/reviewlist"
	"github.com/alex-305/bookbackend/handlers/swagger"
	"github.com/alex-305/bookbackend/handlers/user"
	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(router *mux.Router) {
	//Swagger
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs/swagger/node_modules/swagger-ui-dist/"))))
	router.HandleFunc("/swagger-ui/", swagger.Handle)
	//Auth Routes
	router.HandleFunc("/login", makeHttp(auth.HandleLogin, s.DB)).Methods(http.MethodPost)
	router.HandleFunc("/signup", makeHttp(auth.HandleSignUp, s.DB)).Methods(http.MethodPost)
	//User Routes
	router.HandleFunc("/user/{username}", makeHttp(user.HandleGetUser, s.DB)).Methods(http.MethodGet)
	router.HandleFunc("/user/{username}/description", makeHttp(user.HandlePatchDesc, s.DB)).Methods(http.MethodPatch)
	router.HandleFunc("/user/{username}/password", makeHttp(user.HandlePatchPass, s.DB)).Methods(http.MethodPatch)
	//Review Routes
	router.HandleFunc("/review", makeHttp(review.HandlePost, s.DB)).Methods(http.MethodPost)
	router.HandleFunc("/review/{reviewid}", makeHttp(review.HandlePut, s.DB)).Methods(http.MethodPut)
	router.HandleFunc("/review/{reviewid}", makeHttp(review.HandleDelete, s.DB)).Methods(http.MethodDelete)
	router.HandleFunc("/review/{reviewid}", makeHttp(review.HandleGet, s.DB)).Methods(http.MethodGet)
	//User Review List
	router.HandleFunc("/user/{username}/reviews", makeHttp(reviewlist.HandleGetUser, s.DB)).Methods(http.MethodGet)
	router.HandleFunc("/user/{username}/reviews/count", makeHttp(reviewlist.HandleGetUserCount, s.DB)).Methods(http.MethodGet)
	//Book Review List
	//router.HandleFunc("/reviews/popular", makeHttp(reviewList.HandleGetPopular, s.DB)).Methods(http.MethodGet)
	router.HandleFunc("/book/{volumeid}/reviews", makeHttp(reviewlist.HandleGetBook, s.DB)).Methods(http.MethodGet)
	router.HandleFunc("/book/{volumeid}/reviews/count", makeHttp(reviewlist.HandleGetBookCount, s.DB)).Methods(http.MethodGet)
	//Reply Routes
	router.HandleFunc("/review/{reviewid}/comment", makeHttp(comment.HandlePost, s.DB)).Methods(http.MethodPost)
	router.HandleFunc("/review/{reviewid}/comments", makeHttp(commentlist.HandleGetReview, s.DB)).Methods(http.MethodGet)
	//router.HandleFunc("/comment/{commentid}", makeHttp(comment.HandleGet, s.DB)).Methods(http.MethodGet)
	//Like Routes

}

type handlerFunc func(http.ResponseWriter, *http.Request, *db.DB)
type httpFunc func(http.ResponseWriter, *http.Request)

func makeHttp(fn handlerFunc, d *db.DB) httpFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, d)
	})
}
