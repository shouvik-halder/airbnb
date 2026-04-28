package db

import (
	// "database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImpl struct {
	// sqlDB *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
func (u *UserRepositoryImpl) Create() error {
	fmt.Println("User repository implementation create user")
	return nil
}
