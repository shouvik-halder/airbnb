package dtos

type CreateRoleDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type UpdateRoleDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CreatePermissionDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Resource    string `json:"resource" validate:"required"`
	Action      string `json:"action" validate:"required"`
}

type UpdatePermissionDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Resource    *string `json:"resource"`
	Action      *string `json:"action"`
}

type GetRoleByIdDTO struct {
	Id int64 `json:"id" validate:"required"`
}

type GetRoleByNameDTO struct {
	Name string `json:"name" validate:"required"`
}

type GetPermissionByIdDTO struct {
	Id int64 `json:"id" validate:"required"`
}

type GetPermissionByNameDTO struct {
	Name string `json:"name" validate:"required"`
}

type AssignRoleToUserDTO struct {
	UserId int64 `json:"userId" validate:"required"`
	RoleId int64 `json:"roleId" validate:"required"`
}

type RemoveRoleFromUserDTO struct {
	UserId int64 `json:"userId" validate:"required"`
	RoleId int64 `json:"roleId" validate:"required"`
}

type RolePermissionDTO struct {
	RoleId       int64 `json:"roleId" validate:"required"`
	PermissionId int64 `json:"permissionId" validate:"required"`
}

type UserRolesByUserDTO struct {
	UserId int64 `json:"userId" validate:"required"`
}

type RolePermissionsByRoleDTO struct {
	RoleId int64 `json:"roleId" validate:"required"`
}

type CheckUserPermissionDTO struct {
	UserId         int64  `json:"userId" validate:"required"`
	PermissionName string `json:"permissionName" validate:"required"`
}

type CheckUserRoleDTO struct {
	UserId   int64  `json:"userId" validate:"required"`
	RoleName string `json:"roleName" validate:"required"`
}
