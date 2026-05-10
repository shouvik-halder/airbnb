package services

import (
	db "AuthenticationService/db/repositories"
	"AuthenticationService/model"
)

type UserRolesService interface {
	GetUserRolesByUserIDService(userID int64) ([]*model.UserRole, error)
	AssignUserRoleService(userID, roleID int64) error
	RemoveUserRoleService(userID, roleID int64) error
	HasPermissionService(userID int64, permissionName string) (bool, error)
	HasRoleService(userID int64, roleName string) (bool, error)
}

type userRolesServiceImpl struct {
	userRolesRepo db.UserRolesRepository
}

func NewUserRolesService(userRolesRepo db.UserRolesRepository) UserRolesService {
	return &userRolesServiceImpl{
		userRolesRepo: userRolesRepo,
	}
}

func (s *userRolesServiceImpl) GetUserRolesByUserIDService(userID int64) ([]*model.UserRole, error) {
	return s.userRolesRepo.GetUserRolesByUserID(userID)
}

func (s *userRolesServiceImpl) AssignUserRoleService(userID, roleID int64) error {
	return s.userRolesRepo.AssignUserRole(userID, roleID)
}

func (s *userRolesServiceImpl) RemoveUserRoleService(userID, roleID int64) error {
	return s.userRolesRepo.RemoveUserRole(userID, roleID)
}

func (s *userRolesServiceImpl) HasPermissionService(userID int64, permissionName string) (bool, error) {
	return s.userRolesRepo.HasPermission(userID, permissionName)
}

func (s *userRolesServiceImpl) HasRoleService(userID int64, roleName string) (bool, error) {
	return s.userRolesRepo.HasRole(userID, roleName)
}
