package db

import (
	dbConfig "ReviewService/config/db"
	"ReviewService/repositories"
)

type Storage struct{
	ReviewRepository repositories.ReviewRepository
}

func InitStorage() *Storage{
	return &Storage{
		ReviewRepository: &repositories.ReviewRepositoryImpl{
			SqlDB: dbConfig.GetDB(),
		},
	}
}