package router

import (
	"AuthenticationService/controllers"
	v1router "AuthenticationService/router/v1"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func InitializeRouter(userRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping", controllers.PingController)

	v1Router := v1router.NewV1Router(userRouter)
	v1Router.Register(chiRouter)

	return chiRouter
}
