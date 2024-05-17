package api

import (
	"net/http"

	"golang.org/x/time/rate"
)

func rateLimit(f httpFunc) httpFunc {
	limit := rate.NewLimiter(2, 4)
	return httpFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limit.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		} else {
			f(w, r)
		}
	})
}
