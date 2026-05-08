package routers

import (
	routerHelper "ReviewService/routers/helper"
	v1 "ReviewService/routers/v1"

	"github.com/go-chi/chi/v5"
)

func InitializeRouters(router routerHelper.Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	v1Router := v1.NewV1Router(router)
	chiRouter.Route("/api", func(r chi.Router) {
		v1Router.Register(r)
	})

	return chiRouter
}
