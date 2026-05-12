
INSERT INTO roles (name, description)
VALUES ('admin', 'Administrator role with all permissions')
ON DUPLICATE KEY UPDATE
    description = VALUES(description),
    deleted_at = NULL,
    updated_at = CURRENT_TIMESTAMP;

INSERT INTO `user` (email, password_hash)
VALUES (
    'testuser@email.com',
    'pbkdf2_sha256$210000$ZOeT50zN3LaV253keHej7g$9qewenZwoiOCxFHNipaD6rNLrY6AsR0lUjDbnGtLYfQ'
)
ON DUPLICATE KEY UPDATE
    password_hash = VALUES(password_hash),
    deleted_at = NULL,
    updated_at = CURRENT_TIMESTAMP;

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM `user` u
JOIN roles r ON r.name = 'admin'
WHERE u.email = 'testuser@email.com'
ON DUPLICATE KEY UPDATE
    deleted_at = NULL,
    updated_at = CURRENT_TIMESTAMP;
