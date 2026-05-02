package controllers

import (
	"AuthenticationService/services"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userService services.UserService
}

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController) RegisterController(w http.ResponseWriter, r *http.Request) {
	var payload authRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response, err := uc.userService.Register(payload.Email, payload.Password)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrInvalidInput):
			writeError(w, http.StatusBadRequest, "email and password are required; password must be at least 8 characters")
		case errors.Is(err, services.ErrEmailAlreadyInUse):
			writeError(w, http.StatusConflict, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "failed to register user")
		}
		return
	}

	writeJSON(w, http.StatusCreated, response)
}

func (uc *UserController) LoginController(w http.ResponseWriter, r *http.Request) {
	var payload authRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response, err := uc.userService.Login(payload.Email, payload.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			writeError(w, http.StatusUnauthorized, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to login")
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (uc *UserController) GetUserByIdController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := uc.userService.GetUserByIdService(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to fetch user")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (uc *UserController) DeleteUserByIdController(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	_, err = uc.userService.DeleteUserByIdService(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete user")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) GetAllUsersController(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.GetAllUsersService()
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to fetch all users")
	}

	writeJSON(w, http.StatusOK, users)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Message: message})
}
