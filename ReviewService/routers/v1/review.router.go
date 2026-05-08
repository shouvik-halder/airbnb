package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ReviewRouter struct {
}

func NewReviewRouter() *ReviewRouter {
	return &ReviewRouter{}
}

func (reviewRouter *ReviewRouter) Register(r chi.Router) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
}
