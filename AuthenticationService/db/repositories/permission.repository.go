package db

import (
	"AuthenticationService/model"
	"database/sql"
)

type PermissionRepository interface {
	GetPermissionByID(id int64) (*model.Permission, error)
	GetPermissionByName(name string) (*model.Permission, error)
	GetAllPermissions() ([]*model.Permission, error)
	CreatePermission(
		name string,
		description string,
		resource string,
		action string,
	) (*model.Permission, error)
}

type PermissionRepositoryImpl struct {
	sqlDB *sql.DB
}

func NewPermissionRepository(_sqlDB *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		sqlDB: _sqlDB,
	}
}

func (r *PermissionRepositoryImpl) GetPermissionByID(id int64) (*model.Permission, error) {

	query := `
		SELECT 
			id,
			name,
			description,
			resource,
			action,
			created_at,
			updated_at
		FROM permissions
		WHERE id = ?
		AND deleted_at IS NULL
	`

	row := r.sqlDB.QueryRow(query, id)

	permission := &model.Permission{}

	err := row.Scan(
		&permission.Id,
		&permission.Name,
		&permission.Description,
		&permission.Resource,
		&permission.Action,
		&permission.CreatedAt,
		&permission.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (r *PermissionRepositoryImpl) GetPermissionByName(name string) (*model.Permission, error) {

	query := `
		SELECT 
			id,
			name,
			description,
			resource,
			action,
			created_at,
			updated_at
		FROM permissions
		WHERE name = ?
		AND deleted_at IS NULL
	`

	row := r.sqlDB.QueryRow(query, name)

	permission := &model.Permission{}

	err := row.Scan(
		&permission.Id,
		&permission.Name,
		&permission.Description,
		&permission.Resource,
		&permission.Action,
		&permission.CreatedAt,
		&permission.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (r *PermissionRepositoryImpl) GetAllPermissions() ([]*model.Permission, error) {

	query := `
		SELECT 
			id,
			name,
			description,
			resource,
			action,
			created_at,
			updated_at
		FROM permissions
		WHERE deleted_at IS NULL
	`

	rows, err := r.sqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []*model.Permission{}

	for rows.Next() {

		permission := &model.Permission{}

		err := rows.Scan(
			&permission.Id,
			&permission.Name,
			&permission.Description,
			&permission.Resource,
			&permission.Action,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *PermissionRepositoryImpl) CreatePermission(
	name string,
	description string,
	resource string,
	action string,
) (*model.Permission, error) {

	query := `
		INSERT INTO permissions (
			name,
			description,
			resource,
			action
		)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.sqlDB.Exec(
		query,
		name,
		description,
		resource,
		action,
	)

	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetPermissionByID(insertedID)
}
