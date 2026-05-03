package controllers

import (
	"AuthenticationService/dtos"
	"AuthenticationService/helper"
	"AuthenticationService/services"
	"AuthenticationService/utils"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {

	payload, ok := helper.GetPayLoad[dtos.RegisterRequestDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnprocessableEntity, "invalid json")
		return
	}

	response, err := uc.userService.Register(payload.Email, payload.Password)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrInvalidInput):
			utils.WriteError(w, http.StatusBadRequest, "email and password are required; password must be at least 8 characters")
		case errors.Is(err, services.ErrEmailAlreadyInUse):
			utils.WriteError(w, http.StatusConflict, err.Error())
		default:
			utils.WriteError(w, http.StatusInternalServerError, "failed to register user")
		}
		return
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

func (uc *UserController) LoginController(w http.ResponseWriter, r *http.Request) {

	payload, ok := helper.GetPayLoad[dtos.LoginRequestDTO](r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnprocessableEntity, "invalid json")
		return
	}

	response, err := uc.userService.Login(payload.Email, payload.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			utils.WriteError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "failed to login")
		return
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (uc *UserController) GetUserByIdController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := uc.userService.GetUserByIdService(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "failed to fetch user")
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (uc *UserController) DeleteUserByIdController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	_, err = uc.userService.DeleteUserByIdService(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "failed to delete user")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) GetAllUsersController(w http.ResponseWriter, r *http.Request) {
	log := helper.LoggerFromContext(r.Context())

	log.Info().Msg("getting all users")
	users, err := uc.userService.GetAllUsersService()
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			log.Info().Msg(fmt.Sprintf("error while fetching all active users %s", err.Error()))
			utils.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "failed to fetch all users")
	}

	utils.WriteJSON(w, http.StatusOK, users)
}
