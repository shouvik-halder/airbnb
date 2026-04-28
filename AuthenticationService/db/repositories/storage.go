package db

// "database/sql"

type Storage struct {
	UserRepository UserRepository
}

func InitStorage() *Storage {
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
}
