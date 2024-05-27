package api

import (
	"net/http"

	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/handlers/auth"
	"github.com/alex-305/bookbackend/internal/handlers/comment"
	"github.com/alex-305/bookbackend/internal/handlers/comment/commentlikes"
	"github.com/alex-305/bookbackend/internal/handlers/comment/commentlist"
	"github.com/alex-305/bookbackend/internal/handlers/follow"
	"github.com/alex-305/bookbackend/internal/handlers/review"
	"github.com/alex-305/bookbackend/internal/handlers/review/reviewlikes"
	"github.com/alex-305/bookbackend/internal/handlers/review/reviewlist"
	"github.com/alex-305/bookbackend/internal/handlers/review/reviewlist/reviewliststats"
	"github.com/alex-305/bookbackend/internal/handlers/review/reviewstats"
	"github.com/alex-305/bookbackend/internal/handlers/swagger"
	"github.com/alex-305/bookbackend/internal/handlers/user"
	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(r *mux.Router) {
	//Swagger
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs/swagger/node_modules/swagger-ui-dist/"))))
	r.HandleFunc("/swagger-ui/", swagger.Handle)
	//Auth Routes
	r.HandleFunc("/login", makeHttp(auth.HandleLogin, s.DB)).Methods(http.MethodPost)
	r.HandleFunc("/signup", makeHttp(auth.HandleSignUp, s.DB)).Methods(http.MethodPost)
	//User Routes
	r.HandleFunc("/user/{username}", makeHttp(user.HandleGetUser, s.DB)).Methods(http.MethodGet)
	r.HandleFunc("/user/{username}/description", makeHttp(user.HandlePatchDesc, s.DB)).Methods(http.MethodPatch)
	r.HandleFunc("/user/{username}/password", makeHttp(user.HandlePatchPass, s.DB)).Methods(http.MethodPatch)
	//Review Routes
	r.HandleFunc("/review", makeHttp(review.HandlePost, s.DB)).Methods(http.MethodPost)
	r.HandleFunc("/review/{reviewid}", makeHttp(review.HandlePut, s.DB)).Methods(http.MethodPut)
	r.HandleFunc("/review/{reviewid}", makeHttp(review.HandleDelete, s.DB)).Methods(http.MethodDelete)
	r.HandleFunc("/review/{reviewid}", makeHttp(review.HandleGet, s.DB)).Methods(http.MethodGet)
	//Review Like Routes
	r.HandleFunc("/review/{reviewid}/likes", makeHttp(reviewlikes.HandlePost, s.DB)).Methods(http.MethodPost)
	r.HandleFunc("/review/{reviewid}/likes", makeHttp(reviewlikes.HandleDelete, s.DB)).Methods(http.MethodDelete)
	//User Review List
	r.HandleFunc("/user/{username}/reviews", makeHttp(reviewlist.HandleGetUser, s.DB)).Methods(http.MethodGet)
	r.HandleFunc("/user/{username}/reviews/stats", makeHttp(reviewliststats.HandleGetUser, s.DB)).Methods(http.MethodGet)
	r.HandleFunc("/reviews/following", makeHttp(reviewlist.HandleGetFollowing, s.DB)).Methods(http.MethodGet)
	//Follow Routes
	r.HandleFunc("/user/{username}/follows", makeHttp(follow.HandlePost, s.DB)).Methods(http.MethodPost)
	r.HandleFunc("/user/{username}/follows", makeHttp(follow.HandleDelete, s.DB)).Methods(http.MethodDelete)
	//Follow List
	r.HandleFunc("/user/{username}/following", makeHttp(follow.HandleGetFollowing, s.DB)).Methods(http.MethodGet)
	r.HandleFunc("/user/{username}/followers", makeHttp(follow.HandleGetFollowers, s.DB)).Methods(http.MethodGet)
	//Book Review List
	r.HandleFunc("/volume/{volumeid}/stats", makeHttp(reviewliststats.HandleGetBook, s.DB)).Methods(http.MethodGet)
	r.HandleFunc("/volume/{volumeid}/reviews", makeHttp(reviewlist.HandleGetBook, s.DB)).Methods(http.MethodGet)
	//Book stats list
	//r.HandleFunc("/volumes/popular", makeHttp(reviewList.HandleGetPopular, s.DB)).Methods(http.MethodGet)
	//Comment Routes
	r.HandleFunc("/review/{reviewid}/comments", makeHttp(comment.HandlePost, s.DB)).Methods(http.MethodPost)
	r.HandleFunc("/review/{reviewid}/comments", makeHttp(commentlist.HandleGetReview, s.DB)).Methods(http.MethodGet)
	r.HandleFunc("/review/{reviewid}/stats", makeHttp(reviewstats.HandleGet, s.DB)).Methods(http.MethodGet)
	r.HandleFunc("/comment/{commentid}", makeHttp(comment.HandleDelete, s.DB)).Methods(http.MethodDelete)
	r.HandleFunc("/comment/{commentid}", makeHttp(comment.HandleGet, s.DB)).Methods(http.MethodGet)
	//Comment Like Routes
	r.HandleFunc("/comment/{commentid}/likes", makeHttp(commentlikes.HandlePost, s.DB)).Methods(http.MethodPost)
	r.HandleFunc("/comment/{commentid}/likes", makeHttp(commentlikes.HandleDelete, s.DB)).Methods(http.MethodDelete)

}

type handlerFunc func(http.ResponseWriter, *http.Request, *db.DB)
type httpFunc func(http.ResponseWriter, *http.Request)

func makeHttp(fn handlerFunc, d *db.DB) httpFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, d)
	})
}
