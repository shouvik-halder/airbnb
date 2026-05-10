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
	UpdateRole(id int64, name, description *string) (*model.Role, error)
	DeleteRole(id int64) (bool, error)
	AssignPermissionToRole(roleID, permissionID int64) error
	RemovePermissionFromRole(roleID, permissionID int64) error
	GetPermissionsByRoleID(roleID int64) ([]*model.Permission, error)
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
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

func (r *rolesRepositoryImpl) UpdateRole(id int64, name, description *string) (*model.Role, error) {
	query := `
		UPDATE roles
		SET
			name = COALESCE(?, name),
			description = COALESCE(?, description)
		WHERE id = ?
		AND deleted_at IS NULL
	`

	result, err := r.sqlDB.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		role, err := r.GetRoleById(id)
		if err != nil {
			return nil, err
		}
		if role == nil {
			return nil, sql.ErrNoRows
		}
		return role, nil
	}

	return r.GetRoleById(id)
}

func (r *rolesRepositoryImpl) DeleteRole(id int64) (bool, error) {
	query := `UPDATE roles SET deleted_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL`

	result, err := r.sqlDB.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *rolesRepositoryImpl) AssignPermissionToRole(roleID, permissionID int64) error {
	query := `
		INSERT INTO permission_role (role_id, permission_id, created_at, updated_at)
		VALUES (?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE deleted_at = NULL, updated_at = NOW()
	`

	_, err := r.sqlDB.Exec(query, roleID, permissionID)
	return err
}

func (r *rolesRepositoryImpl) RemovePermissionFromRole(roleID, permissionID int64) error {
	query := `
		UPDATE permission_role
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE role_id = ?
		AND permission_id = ?
		AND deleted_at IS NULL
	`

	_, err := r.sqlDB.Exec(query, roleID, permissionID)
	return err
}

func (r *rolesRepositoryImpl) GetPermissionsByRoleID(roleID int64) ([]*model.Permission, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.description,
			p.resource,
			p.action,
			p.created_at,
			p.updated_at
		FROM permission_role pr
		JOIN permissions p ON pr.permission_id = p.id
		WHERE pr.role_id = ?
		AND pr.deleted_at IS NULL
		AND p.deleted_at IS NULL
	`

	rows, err := r.sqlDB.Query(query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []*model.Permission{}
	for rows.Next() {
		permission := &model.Permission{}
		if err := rows.Scan(
			&permission.Id,
			&permission.Name,
			&permission.Description,
			&permission.Resource,
			&permission.Action,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
