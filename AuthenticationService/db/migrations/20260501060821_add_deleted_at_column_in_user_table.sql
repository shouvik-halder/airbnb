-- +goose Up
ALTER TABLE user
        ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL;

-- +goose Down
ALTER TABLE user
        DROP COLUMN deleted_at;
