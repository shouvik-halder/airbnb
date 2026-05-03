package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/dtos"
	"AuthenticationService/validators"

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
	r.Route("/user", func(r chi.Router) {
		r.With(validators.Validate[dtos.LoginRequestDTO]()).Post("/register", userRouter.userController.RegisterController)
		r.With(validators.Validate[dtos.LoginRequestDTO]()).Post("/login", userRouter.userController.LoginController)
		r.Get("/{id}", userRouter.userController.GetUserByIdController)
		r.Delete("/{id}", userRouter.userController.DeleteUserByIdController)
		r.Get("/", userRouter.userController.GetAllUsersController)
	})

}
