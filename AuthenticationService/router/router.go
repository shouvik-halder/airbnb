package router

import (
	"AuthenticationService/controllers"
	"AuthenticationService/middlewares"
	routerhelper "AuthenticationService/router/helper"
	v1 "AuthenticationService/router/v1"
	"AuthenticationService/utils"

	"github.com/go-chi/chi/v5"
)

func InitializeRouter(router ...routerhelper.Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middlewares.RateLimit)
	chiRouter.Use(middlewares.CorrelationId)
	chiRouter.Use(middlewares.Logger)
	chiRouter.Get("/ping", controllers.PingController)

	chiRouter.HandleFunc("/products", utils.ReverseProxyToService("https://fakestoreapi.com"))

	v1Router := v1.NewV1Router(router...)
	v1Router.Register(chiRouter)

	return chiRouter
}
