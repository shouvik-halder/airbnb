package services

import (
	db "AuthenticationService/db/repositories"
	"AuthenticationService/model"
	"errors"
)

var ErrRoleNotFound = errors.New("role not found")

type RolesService interface {
	GetRoleByIdService(id int64) (*model.Role, error)
	GetRoleByNameService(name string) (*model.Role, error)
	GetAllRolesService() ([]*model.Role, error)
	CreateRoleService(name, description string) (*model.Role, error)
	UpdateRoleService(id int64, name, description *string) (*model.Role, error)
	DeleteRoleService(id int64) error
	AssignPermissionToRoleService(roleID, permissionID int64) error
	RemovePermissionFromRoleService(roleID, permissionID int64) error
	GetPermissionsByRoleIDService(roleID int64) ([]*model.Permission, error)
}

type rolesServiceImpl struct {
	_rolesRepo db.RolesRepository
}

func NewRolesService(rolesRepo db.RolesRepository) RolesService {
	return &rolesServiceImpl{
		_rolesRepo: rolesRepo,
	}
}

func (s *rolesServiceImpl) GetRoleByIdService(id int64) (*model.Role, error) {

	return s._rolesRepo.GetRoleById(id)

}

func (s *rolesServiceImpl) GetRoleByNameService(name string) (*model.Role, error) {
	return s._rolesRepo.GetRoleByName(name)
}

func (s *rolesServiceImpl) GetAllRolesService() ([]*model.Role, error) {
	return s._rolesRepo.GetAllRoles()
}

func (s *rolesServiceImpl) CreateRoleService(name, description string) (*model.Role, error) {
	return s._rolesRepo.CreateRole(name, description)
}

func (s *rolesServiceImpl) UpdateRoleService(id int64, name, description *string) (*model.Role, error) {
	role, err := s._rolesRepo.UpdateRole(id, name, description)
	if err != nil {
		if db.IsNotFound(err) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}
	return role, nil
}

func (s *rolesServiceImpl) DeleteRoleService(id int64) error {
	deleted, err := s._rolesRepo.DeleteRole(id)
	if err != nil {
		return err
	}
	if !deleted {
		return ErrRoleNotFound
	}
	return nil
}

func (s *rolesServiceImpl) AssignPermissionToRoleService(roleID, permissionID int64) error {
	return s._rolesRepo.AssignPermissionToRole(roleID, permissionID)
}

func (s *rolesServiceImpl) RemovePermissionFromRoleService(roleID, permissionID int64) error {
	return s._rolesRepo.RemovePermissionFromRole(roleID, permissionID)
}

func (s *rolesServiceImpl) GetPermissionsByRoleIDService(roleID int64) ([]*model.Permission, error) {
	return s._rolesRepo.GetPermissionsByRoleID(roleID)
}
