-- +goose Up
CREATE TABLE IF NOT EXISTS permission_role (
    id INT AUTO_INCREMENT PRIMARY KEY,
    role_id INT NOT NULL,
    permission_id INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,

    UNIQUE KEY unique_role_permission (role_id, permission_id)
);

-- seeder data
INSERT INTO permission_role (role_id, permission_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 4),
(1, 5), (1, 6), (1, 7), (1, 8),
(1, 9), (1, 10), (1, 11), (1, 12), (1, 13),

(2, 2), (2, 3),

(3, 2), (3, 3), (3, 5), (3, 6);

-- +goose Down
DROP TABLE IF EXISTS permission_role;
