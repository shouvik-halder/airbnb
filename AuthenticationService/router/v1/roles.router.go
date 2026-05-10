package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/dtos"
	"AuthenticationService/middlewares"
	"AuthenticationService/validators"

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
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_read")).
			Get("/", rolesRouter.rolesController.GetRolesController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_read"), validators.ValidateParams[dtos.GetRoleByIdDTO]()).
			Get("/{id}", rolesRouter.rolesController.GetRoleByIdController)

		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_create"), validators.Validate[dtos.CreateRoleDTO]()).
			Post("/", rolesRouter.rolesController.CreateRoleService)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_update"), validators.ValidateParams[dtos.GetRoleByIdDTO](), validators.Validate[dtos.UpdateRoleDTO]()).
			Patch("/{id}", rolesRouter.rolesController.UpdateRoleController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_delete"), validators.ValidateParams[dtos.GetRoleByIdDTO]()).
			Delete("/{id}", rolesRouter.rolesController.DeleteRoleController)

		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_read"), validators.ValidateParams[dtos.RolePermissionsByRoleDTO]()).
			Get("/{roleId}/permissions", rolesRouter.rolesController.GetPermissionsByRoleController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_manage"), validators.ValidateParams[dtos.RolePermissionDTO]()).
			Post("/{roleId}/permissions/{permissionId}", rolesRouter.rolesController.AssignPermissionToRoleController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_manage"), validators.ValidateParams[dtos.RolePermissionDTO]()).
			Delete("/{roleId}/permissions/{permissionId}", rolesRouter.rolesController.RemovePermissionFromRoleController)
	})
}
