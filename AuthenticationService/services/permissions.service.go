package services

import (
	db "AuthenticationService/db/repositories"
	"AuthenticationService/model"
	"errors"
)

type PermissionsService interface {
	GetPermissionIdService(id int64) (*model.Permission, error)
	GetPermissionByNameService(name string) (*model.Permission, error)
	GetAllPermissionService() ([]*model.Permission, error)
	CreatePermissionService(name, description, resource, action string) (*model.Permission, error)
	UpdatePermissionService(id int64, name, description, resource, action *string) (*model.Permission, error)
	DeletePermissionService(id int64) error
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

func (s *permissionsServiceImpl) UpdatePermissionService(id int64, name, description, resource, action *string) (*model.Permission, error) {
	permission, err := s._permissionsRepo.UpdatePermission(id, name, description, resource, action)
	if err != nil {
		if db.IsNotFound(err) {
			return nil, ErrPermissionNotFound
		}
		return nil, err
	}
	return permission, nil
}

func (s *permissionsServiceImpl) DeletePermissionService(id int64) error {
	deleted, err := s._permissionsRepo.DeletePermission(id)
	if err != nil {
		return err
	}
	if !deleted {
		return ErrPermissionNotFound
	}
	return nil
}

var ErrPermissionNotFound = errors.New("permission not found")
