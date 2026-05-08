package v1

import (
	"ReviewService/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ReviewRouter struct {
	reviewController *controllers.ReviewController
}

func NewReviewRouter(_reviewController *controllers.ReviewController) *ReviewRouter {
	return &ReviewRouter{
		reviewController: _reviewController,
	}
}

func (reviewRouter *ReviewRouter) Register(r chi.Router) {
	r.Route("/reviews", func(r chi.Router) {

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		r.Get("/", reviewRouter.reviewController.CreateReviewController)
	})
}
