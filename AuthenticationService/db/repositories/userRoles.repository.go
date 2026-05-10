package db

import (
	"AuthenticationService/model"
	"database/sql"
)

type UserRolesRepository interface {
	GetUserRolesByUserID(userID int64) ([]*model.UserRole, error)
	AssignUserRole(userID int64, roleID int64) error
	RemoveUserRole(userID int64, roleID int64) error
	HasPermission(userID int64, permissionName string) (bool, error)
	HasRole(userID int64, roleName string) (bool, error)
}

type userRolesrepositoryImpl struct {
	sqlDB *sql.DB
}

func NewUserRolesRepository(sqlDB *sql.DB) UserRolesRepository {
	return &userRolesrepositoryImpl{
		sqlDB: sqlDB,
	}
}

func (r *userRolesrepositoryImpl) GetUserRolesByUserID(userID int64) ([]*model.UserRole, error) {

	query := `
		SELECT 
			ur.id,
			ur.user_id,
			ur.role_id,
			ur.created_at,
			ur.updated_at
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ?
		AND ur.deleted_at IS NULL
		AND r.deleted_at IS NULL
	`

	rows, err := r.sqlDB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userRoles := []*model.UserRole{}

	for rows.Next() {
		userRole := &model.UserRole{}

		err := rows.Scan(
			&userRole.Id,
			&userRole.UserId,
			&userRole.RoleId,
			&userRole.CreatedAt,
			&userRole.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, userRole)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userRoles, nil
}

func (r *userRolesrepositoryImpl) AssignUserRole(userID int64, roleID int64) error {
	query := `
		INSERT INTO user_roles (user_id, role_id, created_at, updated_at)
		VALUES (?, ?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE deleted_at = NULL, updated_at = NOW()
	`

	_, err := r.sqlDB.Exec(query, userID, roleID)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRolesrepositoryImpl) RemoveUserRole(userID int64, roleID int64) error {
	query := `
		UPDATE user_roles
		SET deleted_at = NOW(),
		    updated_at = NOW()
		WHERE user_id = ? 
		  AND role_id = ?
		  AND deleted_at IS NULL
	`

	_, err := r.sqlDB.Exec(query, userID, roleID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRolesrepositoryImpl) HasPermission(userID int64, permissionName string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM user_roles ur
			JOIN roles r 
				ON ur.role_id = r.id
			JOIN permission_role rp 
				ON r.id = rp.role_id
			JOIN permissions p 
				ON rp.permission_id = p.id
			WHERE ur.user_id = ?
				AND p.name = ?
				AND ur.deleted_at IS NULL
				AND r.deleted_at IS NULL
				AND rp.deleted_at IS NULL
				AND p.deleted_at IS NULL
		)
	`

	var exists bool

	err := r.sqlDB.QueryRow(query, userID, permissionName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRolesrepositoryImpl) HasRole(userID int64, roleName string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM user_roles ur
			JOIN roles r 
				ON ur.role_id = r.id
			WHERE ur.user_id = ?
				AND r.name = ?
				AND ur.deleted_at IS NULL
				AND r.deleted_at IS NULL
		)
	`

	var exists bool
	err := r.sqlDB.QueryRow(query, userID, roleName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
