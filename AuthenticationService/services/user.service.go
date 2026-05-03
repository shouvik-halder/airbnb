package services

import (
	"AuthenticationService/config"
	db "AuthenticationService/db/repositories"
	"AuthenticationService/model"
	"AuthenticationService/utils"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidInput        = errors.New("invalid input")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrEmailAlreadyInUse   = errors.New("email already in use")
	ErrUserNotFound        = errors.New("user not found")
	ErrTokenSecretRequired = errors.New("token secret is required")
)

type UserService interface {
	Register(email, password string) (*AuthResponse, error)
	Login(email, password string) (*AuthResponse, error)
	GetUserByIdService(id int64) (*model.User, error)
	DeleteUserByIdService(id int64) (bool, error)
	GetAllUsersService() ([]*model.User, error)
}

type AuthResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

type userServiceImpl struct {
	userRepository db.UserRepository
	tokenSecret    string
}

func (userService *userServiceImpl) Register(email, password string) (*AuthResponse, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || !strings.Contains(email, "@") || len(password) < 8 {
		return nil, ErrInvalidInput
	}
	_, err := userService.userRepository.GetByEmail(email)
	if err == nil {
		return nil, ErrEmailAlreadyInUse
	}
	if !db.IsNotFound(err) {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := userService.userRepository.Create(email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return userService.buildAuthResponse(user)
}

func (userService *userServiceImpl) Login(email, password string) (*AuthResponse, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	user, err := userService.userRepository.GetByEmail(email)
	if err != nil {
		if db.IsNotFound(err) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	ok, err := utils.VerifyPassword(password, user.PasswordHash)
	if err != nil || !ok {
		return nil, ErrInvalidCredentials
	}

	return userService.buildAuthResponse(user)
}

func (userService *userServiceImpl) GetUserByIdService(id int64) (*model.User, error) {
	user, err := userService.userRepository.GetById(id)
	if err != nil {
		if db.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (userservice *userServiceImpl) GetAllUsersService() ([]*model.User, error) {
	users, err := userservice.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	if db.IsNotFound(err) {
		return nil, ErrUserNotFound
	}
	return users, nil
}

func (userService *userServiceImpl) DeleteUserByIdService(id int64) (bool, error) {
	isDeleted, err := userService.userRepository.DeleteById(id)
	if err != nil {
		return false, err
	}

	if !isDeleted {
		return false, ErrUserNotFound
	}

	return true, nil
}

func (userService *userServiceImpl) buildAuthResponse(user *model.User) (*AuthResponse, error) {
	token, err := userService.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func (userService *userServiceImpl) generateToken(user *model.User) (string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(user.Id, 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "Airbnb-Authentication-Service",
	}
	tokenSecret := config.GetConfig().Auth.TokenSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenSecret))
}

func NewUserService(_userRepository db.UserRepository, tokenSecret string) UserService {
	return &userServiceImpl{
		userRepository: _userRepository,
		tokenSecret:    tokenSecret,
	}
}
