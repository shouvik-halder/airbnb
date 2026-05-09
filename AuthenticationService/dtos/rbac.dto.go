package dtos

type CreateRoleDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CreatePermissionDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
}
