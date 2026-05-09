package db

import (
	"AuthenticationService/model"
	"database/sql"
)

type RolesRepository interface {
	GetRoleById(id int64) (*model.Role, error)
	GetRoleByName(name string) (*model.Role, error)
	GetAllRoles() ([]*model.Role, error)
	CreateRole(name, description string) (*model.Role, error)
}

type rolesRepositoryImpl struct {
	sqlDB *sql.DB
}

func NewRolesRepository(_sqlDB *sql.DB) RolesRepository {
	return &rolesRepositoryImpl{
		sqlDB: _sqlDB,
	}
}

func (r *rolesRepositoryImpl) GetRoleById(id int64) (*model.Role, error) {

	query := `SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ? AND deleted_at IS NULL`

	row := r.sqlDB.QueryRow(query, id)
	role := &model.Role{}
	if err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
		return nil, err
	}

	return role, nil
}

func (r *rolesRepositoryImpl) GetRoleByName(name string) (*model.Role, error) {

	query := `SELECT id, name, description, created_at, updated_at 
			  FROM roles 
			  WHERE name = ? AND deleted_at IS NULL`

	row := r.sqlDB.QueryRow(query, name)

	role := &model.Role{}

	if err := row.Scan(
		&role.Id,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return role, nil
}

func (r *rolesRepositoryImpl) GetAllRoles() ([]*model.Role, error) {

	query := `
		SELECT id, name, description, created_at, updated_at
		FROM roles
		WHERE deleted_at IS NULL
	`

	rows, err := r.sqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*model.Role{}

	for rows.Next() {

		role := &model.Role{}

		err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *rolesRepositoryImpl) CreateRole(name, description string) (*model.Role, error) {

	query := `
		INSERT INTO roles (name, description)
		VALUES (?, ?)
	`

	result, err := r.sqlDB.Exec(query, name, description)
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetRoleById(insertedID)
}
