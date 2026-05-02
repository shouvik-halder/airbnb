package db

import (
	"AuthenticationService/model"
	"database/sql"
)

type UserRepository interface {
	GetById(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(email, hashedPass string) (*model.User, error)
	DeleteById(id int64) (bool, error)
	GetAllUsers() ([]*model.User, error)
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
	query := `SELECT id, email, password_hash, created_at, updated_at FROM user WHERE id = ? AND deleted_at IS NULL`
	row := u.sqlDB.QueryRow(query, id)

	user := &model.User{}

	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, password_hash, created_at, updated_at FROM user WHERE email = ? AND deleted_at IS NULL`
	row := u.sqlDB.QueryRow(query, email)

	user := &model.User{}

	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Create(email, hashedPass string) (*model.User, error) {
	query := `INSERT INTO user (email, password_hash) values (?, ?)`
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

func (u *UserRepositoryImpl) DeleteById(id int64) (bool, error) {
	query := `UPDATE user SET deleted_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL`
	result, err := u.sqlDB.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (u *UserRepositoryImpl) GetAllUsers() ([]*model.User, error) {
	query := `SELECT id, email, password_hash, created_at, updated_at FROM user WHERE deleted_at IS NULL`
	row, err := u.sqlDB.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*model.User{}

	for row.Next() {
		var user model.User
		err := row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func IsNotFound(err error) bool {
	return err == sql.ErrNoRows
}
