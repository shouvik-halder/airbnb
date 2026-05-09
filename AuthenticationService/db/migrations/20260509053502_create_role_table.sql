-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- seeder data
INSERT INTO roles (name, description) VALUES
('admin', 'Administrator role with all permissions'),
('user', 'Regular user role with limited permissions'),
('manager', 'Manager role with elevated permissions');

-- +goose Down
DROP TABLE IF EXISTS roles;