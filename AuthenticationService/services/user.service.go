package services

import (
	db "AuthenticationService/db/repositories"
	"AuthenticationService/model"
	"crypto/hmac"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
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

	hashedPassword, err := hashPassword(password)
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

	ok, err := verifyPassword(password, user.PasswordHash)
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

func (userService *userServiceImpl) DeleteUserByIdService(id int64) (bool, error) {
	isDeleted, err := userService.userRepository.Delete(id)
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
	if strings.TrimSpace(userService.tokenSecret) == "" {
		return "", ErrTokenSecretRequired
	}

	now := time.Now()
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	payload := map[string]any{
		"sub":   user.Id,
		"email": user.Email,
		"exp":   now.Add(24 * time.Hour).Unix(),
		"iat":   now.Unix(),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	unsignedToken := fmt.Sprintf("%s.%s", base64.RawURLEncoding.EncodeToString(headerJSON), base64.RawURLEncoding.EncodeToString(payloadJSON))
	mac := hmac.New(sha256.New, []byte(userService.tokenSecret))
	mac.Write([]byte(unsignedToken))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s.%s", unsignedToken, signature), nil
}

func hashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	key, err := pbkdf2.Key(sha256.New, password, salt, 210000, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"pbkdf2_sha256$210000$%s$%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key),
	), nil
}

func verifyPassword(password, storedHash string) (bool, error) {
	parts := strings.Split(storedHash, "$")
	if len(parts) != 4 || parts[0] != "pbkdf2_sha256" {
		return false, nil
	}

	var iterations int
	if _, err := fmt.Sscanf(parts[1], "%d", &iterations); err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[2])
	if err != nil {
		return false, err
	}

	expectedKey, err := base64.RawStdEncoding.DecodeString(parts[3])
	if err != nil {
		return false, err
	}

	actualKey, err := pbkdf2.Key(sha256.New, password, salt, iterations, len(expectedKey))
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(actualKey, expectedKey) == 1, nil
}

func NewUserService(_userRepository db.UserRepository, tokenSecret string) UserService {
	return &userServiceImpl{
		userRepository: _userRepository,
		tokenSecret:    tokenSecret,
	}
}
