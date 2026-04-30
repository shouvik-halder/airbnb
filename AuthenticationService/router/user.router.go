package router

import (
	"AuthenticationService/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) *UserRouter {
	return &UserRouter{
		userController: _userController,
	}
}

func (userRouter *UserRouter) Register(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", userRouter.userController.RegisterController)
		r.Post("/login", userRouter.userController.LoginController)
	})

	r.Get("/users/{id}", userRouter.userController.GetUserByIdController)
}
