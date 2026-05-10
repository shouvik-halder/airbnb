package controllers

import (
	"AuthenticationService/dtos"
	"AuthenticationService/helper"
	"AuthenticationService/model"
	"AuthenticationService/services"
	"AuthenticationService/utils"
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
	payload, ok := helper.GetPayLoad[dtos.GetRoleByIdDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with payload")
		return
	}

	result, err := rc.roleService.GetRoleByIdService(payload.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
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
			utils.WriteJSON(w, http.StatusNoContent, []*model.Role{})
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
		utils.WriteJSON(w, http.StatusNoContent, []*model.Role{})
	}

	utils.WriteJSON(w, http.StatusOK, result)
}
