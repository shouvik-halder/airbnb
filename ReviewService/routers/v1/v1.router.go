package v1

import (
	routerHelper "ReviewService/routers/helper"

	"github.com/go-chi/chi/v5"
)


type V1Router struct {
	routers  routerHelper.Router
}

func NewV1Router(routers routerHelper.Router) *V1Router{
	return &V1Router{
		routers: routers,
	}
}

func (v1Router *V1Router) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		v1Router.Register(r)
		// for _, route := range v1Router.routers{
		// }
	})
	
}