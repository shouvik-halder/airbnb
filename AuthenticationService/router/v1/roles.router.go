package v1

import (
	"AuthenticationService/controllers"
	// "AuthenticationService/dtos"
	// "AuthenticationService/validators"

	"github.com/go-chi/chi/v5"
)

type RolesRouter struct {
	rolesController *controllers.RoleController
}

func NewRolesRouter(_rolesController *controllers.RoleController) *RolesRouter {
	return &RolesRouter{
		rolesController: _rolesController,
	}
}

func (rolesRouter *RolesRouter) Register(r chi.Router) {
	r.Route("/roles", func(r chi.Router) {
		r.Get("/", rolesRouter.rolesController.GetAllRolesController);
		// r.With(validators.Validate[dtos.RegisterRequestDTO]()).Post("/register", userRouter.userController.RegisterController)
		// r.With(validators.Validate[dtos.LoginRequestDTO]()).Post("/login", userRouter.userController.LoginController)
		// r.With(middlewares.JWTAuthenticate).Get("/{id}", userRouter.userController.GetUserByIdController)
		// r.With(middlewares.JWTAuthenticate).Delete("/{id}", userRouter.userController.DeleteUserByIdController)
		// r.With(middlewares.JWTAuthenticate).Get("/", userRouter.userController.GetAllUsersController)
	})

}
