package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/dtos"
	"AuthenticationService/middlewares"
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
		r.With(validators.Validate[dtos.RegisterRequestDTO]()).Post("/register", userRouter.userController.RegisterController)
		r.With(validators.Validate[dtos.LoginRequestDTO]()).Post("/login", userRouter.userController.LoginController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequireRole("user")).Get("/{id}", userRouter.userController.GetUserByIdController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("user_delete")).Delete("/{id}", userRouter.userController.DeleteUserByIdController)
		r.With(middlewares.JWTAuthenticate).Get("/", userRouter.userController.GetAllUsersController)
	})

}
