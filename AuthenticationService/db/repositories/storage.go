package db

import (
	dbconfig "AuthenticationService/config/db"
)

type Storage struct {
	UserRepository UserRepository
}

func InitStorage() *Storage {

	return &Storage{
		UserRepository: &UserRepositoryImpl{
			sqlDB: dbconfig.GetDB(),
		},
	}
}
