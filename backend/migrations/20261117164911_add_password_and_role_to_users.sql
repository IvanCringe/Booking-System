-- +goose Up
ALTER TABLE users
    ADD COLUMN password VARCHAR(255) NOT NULL DEFAULT '',
    ADD COLUMN role VARCHAR(50) NOT NULL DEFAULT 'CLIENT'
        CHECK (role IN ('ADMIN', 'OWNER', 'CLIENT'));
-- +goose Down
ALTER TABLE users
    DROP COLUMN password,
    DROP COLUMN role;
