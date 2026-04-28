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

func (urout *UserRouter) Register(r chi.Router) {
	r.Post("/signup", urout.userController.RegisterUser)
}
