package controllers

import (
	"AuthenticationService/dtos"
	"AuthenticationService/helper"
	"AuthenticationService/services"
	"AuthenticationService/utils"
	"errors"
	"net/http"
)

type PermissionController struct {
	permissionService services.PermissionsService
}

func NewPermissionController(permissionService services.PermissionsService) *PermissionController {
	return &PermissionController{
		permissionService: permissionService,
	}
}

func (pc *PermissionController) GetPermissionsController(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		permission, err := pc.permissionService.GetPermissionByNameService(name)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if permission == nil {
			utils.WriteJSON(w, http.StatusOK, []any{})
			return
		}
		utils.WriteJSON(w, http.StatusOK, permission)
		return
	}

	permissions, err := pc.permissionService.GetAllPermissionService()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, permissions)
}

func (pc *PermissionController) GetPermissionByIDController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.GetPermissionByIdDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	permission, err := pc.permissionService.GetPermissionIdService(payload.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if permission == nil {
		utils.WriteError(w, http.StatusNotFound, "permission not found")
		return
	}

	utils.WriteJSON(w, http.StatusOK, permission)
}

func (pc *PermissionController) CreatePermissionController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetPayLoad[dtos.CreatePermissionDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with payload")
		return
	}

	permission, err := pc.permissionService.CreatePermissionService(
		payload.Name,
		payload.Description,
		payload.Resource,
		payload.Action,
	)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusCreated, permission)
}

func (pc *PermissionController) UpdatePermissionController(w http.ResponseWriter, r *http.Request) {
	params, ok := helper.GetParams[dtos.GetPermissionByIdDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	payload, ok := helper.GetPayLoad[dtos.UpdatePermissionDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with payload")
		return
	}

	permission, err := pc.permissionService.UpdatePermissionService(
		params.Id,
		payload.Name,
		payload.Description,
		payload.Resource,
		payload.Action,
	)
	if err != nil {
		if errors.Is(err, services.ErrPermissionNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, permission)
}

func (pc *PermissionController) DeletePermissionController(w http.ResponseWriter, r *http.Request) {
	payload, ok := helper.GetParams[dtos.GetPermissionByIdDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "issue with params")
		return
	}

	if err := pc.permissionService.DeletePermissionService(payload.Id); err != nil {
		if errors.Is(err, services.ErrPermissionNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
