package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/dtos"
	"AuthenticationService/validators"

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
		r.Get("/", rolesRouter.rolesController.GetRolesController)
		r.With(validators.ValidateParams[dtos.GetRoleByIdDTO]()).Get("/{id}", rolesRouter.rolesController.GetRoleByIdController)

		r.With(validators.Validate[dtos.CreateRoleDTO]()).
			Post("/", rolesRouter.rolesController.CreateRoleService)
	})
}
