package services

import (
	db "AuthenticationService/db/repositories"
	"fmt"
)


type UserService interface {
	CreateUser() error
}

type userServiceImpl struct {
	userRepository db.UserRepository
}

func (ur *userServiceImpl) CreateUser() error {
	fmt.Println("User service implementation create user");
	ur.userRepository.Create()
	return  nil;
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: _userRepository,
	}
}