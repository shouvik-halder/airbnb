package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/dtos"
	"AuthenticationService/middlewares"
	"AuthenticationService/validators"

	"github.com/go-chi/chi/v5"
)

type UserRolesRouter struct {
	userRoleController *controllers.UserRoleController
}

func NewUserRolesRouter(userRoleController *controllers.UserRoleController) *UserRolesRouter {
	return &UserRolesRouter{
		userRoleController: userRoleController,
	}
}

func (userRolesRouter *UserRolesRouter) Register(r chi.Router) {
	r.Route("/users/{userId}", func(r chi.Router) {
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_read"), validators.ValidateParams[dtos.UserRolesByUserDTO]()).
			Get("/roles", userRolesRouter.userRoleController.GetUserRolesController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_update"), validators.ValidateParams[dtos.AssignRoleToUserDTO]()).
			Post("/roles/{roleId}", userRolesRouter.userRoleController.AssignUserRoleController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_update"), validators.ValidateParams[dtos.RemoveRoleFromUserDTO]()).
			Delete("/roles/{roleId}", userRolesRouter.userRoleController.RemoveUserRoleController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("role_read"), validators.ValidateParams[dtos.CheckUserRoleDTO]()).
			Get("/roles/{roleName}", userRolesRouter.userRoleController.CheckUserRoleController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_read"), validators.ValidateParams[dtos.CheckUserPermissionDTO]()).
			Get("/permissions/{permissionName}", userRolesRouter.userRoleController.CheckUserPermissionController)
	})
}
