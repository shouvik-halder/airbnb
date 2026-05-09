-- +goose Up
CREATE TABLE IF NOT EXISTS permissions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) UNIQUE,
    description TEXT,
    resource VARCHAR(256) NOT NULL,
    action VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- seeder data
INSERT INTO permissions (name, description, resource, action) VALUES
('user_create', 'Permission to create a user', 'user', 'create'),
('user_read', 'Permission to read user information', 'user', 'read'),
('user_update', 'Permission to update user information', 'user', 'update'),
('user_delete', 'Permission to delete a user', 'user', 'delete'),
('role_create', 'Permission to create a role', 'role', 'create'),
('role_read', 'Permission to read role information', 'role', 'read'),
('role_update', 'Permission to update role information', 'role', 'update'),
('role_delete', 'Permission to delete a role', 'role', 'delete'),
('permission_manage', 'Permission to manage permissions', 'permission', 'manage'),
('permission_create', 'Permission to create a permission', 'permission', 'create'),
('permission_read', 'Permission to read permission information', 'permission', 'read'),
('permission_update', 'Permission to update permission information', 'permission', 'update'),
('permission_delete', 'Permission to delete a permission', 'permission', 'delete');

-- +goose Down
DROP TABLE IF EXISTS permissions;