package controllers

import (
	"AuthenticationService/dtos"
	"AuthenticationService/helper"
	"AuthenticationService/model"
	"AuthenticationService/services"
	"AuthenticationService/utils"
	"errors"
	"net/http"
)

type RoleController struct {
	roleService services.RolesService
}

func NewRoleController(roleService services.RolesService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (rc *RoleController) GetRoleByIdController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.GetRoleByIdDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with payload")
		return
	}

	result, err := rc.roleService.GetRoleByIdService(payload.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if result == nil {
		utils.WriteError(w, http.StatusNotFound, "role not found")
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}

func (rc *RoleController) CreateRoleService(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetPayLoad[dtos.CreateRoleDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with payload")
		return
	}

	response, err := rc.roleService.CreateRoleService(payload.Name, payload.Description)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

func (rc *RoleController) GetRolesController(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// Filter by name
	if name != "" {
		result, err := rc.roleService.GetRoleByNameService(name)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if result == nil {
			utils.WriteJSON(w, http.StatusOK, []*model.Role{})
			return
		}

		utils.WriteJSON(w, http.StatusOK, result)
		return
	}

	// Get all roles
	result, err := rc.roleService.GetAllRolesService()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(result) == 0 {
		utils.WriteJSON(w, http.StatusOK, []*model.Role{})
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}

func (rc *RoleController) UpdateRoleController(w http.ResponseWriter, r *http.Request) {
	params, ok := helper.GetParams[dtos.GetRoleByIdDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	payload, ok := helper.GetPayLoad[dtos.UpdateRoleDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with payload")
		return
	}

	response, err := rc.roleService.UpdateRoleService(params.Id, payload.Name, payload.Description)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (rc *RoleController) DeleteRoleController(w http.ResponseWriter, r *http.Request) {
	params, ok := helper.GetParams[dtos.GetRoleByIdDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	if err := rc.roleService.DeleteRoleService(params.Id); err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rc *RoleController) AssignPermissionToRoleController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.RolePermissionDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	if err := rc.roleService.AssignPermissionToRoleService(payload.RoleId, payload.PermissionId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"message":      "permission assigned to role",
		"roleId":       payload.RoleId,
		"permissionId": payload.PermissionId,
	})
}

func (rc *RoleController) RemovePermissionFromRoleController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.RolePermissionDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	if err := rc.roleService.RemovePermissionFromRoleService(payload.RoleId, payload.PermissionId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rc *RoleController) GetPermissionsByRoleController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.RolePermissionsByRoleDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	permissions, err := rc.roleService.GetPermissionsByRoleIDService(payload.RoleId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, permissions)
}
