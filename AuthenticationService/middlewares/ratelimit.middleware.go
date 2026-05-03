package middlewares

import (
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

func RateLimit(next http.Handler) http.Handler {
	limiter := httprate.LimitByIP(5, time.Minute)

	return limiter(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}))
}