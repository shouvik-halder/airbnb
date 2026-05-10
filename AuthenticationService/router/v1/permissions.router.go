package v1

import (
	"AuthenticationService/controllers"
	"AuthenticationService/dtos"
	"AuthenticationService/middlewares"
	"AuthenticationService/validators"

	"github.com/go-chi/chi/v5"
)

type PermissionsRouter struct {
	permissionController *controllers.PermissionController
}

func NewPermissionsRouter(permissionController *controllers.PermissionController) *PermissionsRouter {
	return &PermissionsRouter{
		permissionController: permissionController,
	}
}

func (permissionsRouter *PermissionsRouter) Register(r chi.Router) {
	r.Route("/permissions", func(r chi.Router) {
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_read")).
			Get("/", permissionsRouter.permissionController.GetPermissionsController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_read"), validators.ValidateParams[dtos.GetPermissionByIdDTO]()).
			Get("/{id}", permissionsRouter.permissionController.GetPermissionByIDController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_create"), validators.Validate[dtos.CreatePermissionDTO]()).
			Post("/", permissionsRouter.permissionController.CreatePermissionController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_update"), validators.ValidateParams[dtos.GetPermissionByIdDTO](), validators.Validate[dtos.UpdatePermissionDTO]()).
			Patch("/{id}", permissionsRouter.permissionController.UpdatePermissionController)
		r.With(middlewares.JWTAuthenticate, middlewares.RequirePermission("permission_delete"), validators.ValidateParams[dtos.GetPermissionByIdDTO]()).
			Delete("/{id}", permissionsRouter.permissionController.DeletePermissionController)
	})
}
