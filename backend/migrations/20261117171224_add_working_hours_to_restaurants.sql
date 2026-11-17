-- +goose Up
ALTER TABLE restaurants
    ADD COLUMN open_time TIME NOT NULL DEFAULT '09:00:00',
    ADD COLUMN close_time TIME NOT NULL DEFAULT '23:00:00';

-- +goose Down
ALTER TABLE restaurants
    DROP COLUMN open_time,
    DROP COLUMN close_time;
