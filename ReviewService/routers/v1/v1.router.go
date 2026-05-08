package v1

import (
	routerHelper "ReviewService/routers/helper"

	"github.com/go-chi/chi/v5"
)

type V1Router struct {
	router routerHelper.Router
}

func NewV1Router(router routerHelper.Router) *V1Router {
	return &V1Router{
		router: router,
	}
}

func (v1Router *V1Router) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {

			v1Router.router.Register(r)

	})
}
