package v1

import (
	routerhelper "AuthenticationService/router/helper"

	"github.com/go-chi/chi/v5"
)



type V1Router struct {
	routers []routerhelper.Router
}

func NewV1Router(routers ...routerhelper.Router) *V1Router {
	return &V1Router{
		routers: routers,
	}
}

func (v1 *V1Router) Register(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		for _, router := range v1.routers {
			router.Register(r)
		}
	})
}
