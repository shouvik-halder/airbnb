package controllers

import (
	"AuthenticationService/services"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	uc.userService.CreateUser()
	w.Write([]byte("User registration endpoint"))
}
