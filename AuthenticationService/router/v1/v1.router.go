package v1

import "github.com/go-chi/chi/v5"

type Router interface {
	Register(r chi.Router)
}

type V1Router struct {
	routers []Router
}

func NewV1Router(routers ...Router) *V1Router {
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
