package db

import (
	"AuthenticationService/model"
	"database/sql"
)

type UserRepository interface {
	GetById(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(email, hashedPass string) (*model.User, error)
}

type UserRepositoryImpl struct {
	sqlDB *sql.DB
}

func NewUserRepository(_sqlDB *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		sqlDB: _sqlDB,
	}
}
func (u *UserRepositoryImpl) GetById(id int64) (*model.User, error) {
	query := `SELECT id, email, password_hash, created_at, updated_at FROM users WHERE id = ?`
	row := u.sqlDB.QueryRow(query, id)

	user := &model.User{}

	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = ?`
	row := u.sqlDB.QueryRow(query, email)

	user := &model.User{}

	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Create(email, hashedPass string) (*model.User, error) {
	query := `INSERT INTO users(email, password_hash) values (?, ?)`
	result, err := u.sqlDB.Exec(query, email, hashedPass)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return u.GetById(id)
}

func IsNotFound(err error) bool {
	return err == sql.ErrNoRows
}
