package services

import (
	db "AuthenticationService/db/repositories"
	"AuthenticationService/model"
)

type RolesService interface {
	GetRoleByIdService(id int64) (*model.Role, error)
	GetRoleByNameService(name string) (*model.Role, error)
	GetAllRolesService() ([]*model.Role, error)
	CreateRoleService(name, description string) (*model.Role, error)
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

