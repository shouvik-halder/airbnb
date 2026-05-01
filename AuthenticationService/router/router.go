package router

import (
	"AuthenticationService/controllers"
	routerhelper "AuthenticationService/router/helper"
	v1 "AuthenticationService/router/v1"

	"github.com/go-chi/chi/v5"
)



func InitializeRouter(router ...routerhelper.Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping", controllers.PingController)

	v1Router := v1.NewV1Router(router...)
	v1Router.Register(chiRouter)

	return chiRouter
}
