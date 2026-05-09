package services

import (
	db "AuthenticationService/db/repositories"
	"AuthenticationService/model"
)

type PermissionsService interface {
	GetPermissionIdService(id int64) (*model.Permission, error)
	GetPermissionByNameService(name string) (*model.Permission, error)
	GetAllPermissionService() ([]*model.Permission, error)
	CreatePermissionService(name, description, resource, action string) (*model.Permission, error)
}

type permissionsServiceImpl struct {
	_permissionsRepo db.PermissionRepository
}

func NewPermissionsService(permissionsRepo db.PermissionRepository) PermissionsService {
	return &permissionsServiceImpl{
		_permissionsRepo: permissionsRepo,
	}
}

func (s *permissionsServiceImpl) GetPermissionIdService(id int64) (*model.Permission, error) {

	return s._permissionsRepo.GetPermissionByID(id)

}

func (s *permissionsServiceImpl) GetPermissionByNameService(name string) (*model.Permission, error) {
	return s._permissionsRepo.GetPermissionByName(name)
}

func (s *permissionsServiceImpl) GetAllPermissionService() ([]*model.Permission, error) {
	return s._permissionsRepo.GetAllPermissions()
}

func (s *permissionsServiceImpl) CreatePermissionService(name, description, resource, action string) (*model.Permission, error) {
	return s._permissionsRepo.CreatePermission(name, description, resource, action)
}
