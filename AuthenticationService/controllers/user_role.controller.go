package controllers

import (
	"AuthenticationService/dtos"
	"AuthenticationService/helper"
	"AuthenticationService/services"
	"AuthenticationService/utils"
	"net/http"
)

type UserRoleController struct {
	userRolesService services.UserRolesService
}

func NewUserRoleController(userRolesService services.UserRolesService) *UserRoleController {
	return &UserRoleController{
		userRolesService: userRolesService,
	}
}

func (uc *UserRoleController) GetUserRolesController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.UserRolesByUserDTO](r.Context())
	logger := helper.LoggerFromContext(r.Context())
	logger.Info().Msgf("payload in get user roles controller is %v", payload)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	roles, err := uc.userRolesService.GetUserRolesByUserIDService(payload.UserId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, roles)
}

func (uc *UserRoleController) AssignUserRoleController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.AssignRoleToUserDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	if err := uc.userRolesService.AssignUserRoleService(payload.UserId, payload.RoleId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"message": "role assigned to user",
		"userId":  payload.UserId,
		"roleId":  payload.RoleId,
	})
}

func (uc *UserRoleController) RemoveUserRoleController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.RemoveRoleFromUserDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	if err := uc.userRolesService.RemoveUserRoleService(payload.UserId, payload.RoleId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserRoleController) CheckUserPermissionController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.CheckUserPermissionDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	hasPermission, err := uc.userRolesService.HasPermissionService(payload.UserId, payload.PermissionName)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"userId":         payload.UserId,
		"permissionName": payload.PermissionName,
		"hasPermission":  hasPermission,
	})
}

func (uc *UserRoleController) CheckUserRoleController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.CheckUserRoleDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	hasRole, err := uc.userRolesService.HasRoleService(payload.UserId, payload.RoleName)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"userId":   payload.UserId,
		"roleName": payload.RoleName,
		"hasRole":  hasRole,
	})
}
