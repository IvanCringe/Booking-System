-- +goose Up
CREATE TABLE restaurant_tables (
    id BIGSERIAL PRIMARY KEY,
    restaurant_id BIGINT NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL, -- Например: 'Стол №1', 'У окна', 'VIP'
    seats_count INT NOT NULL CHECK (seats_count > 0),
    is_active BOOLEAN NOT NULL DEFAULT TRUE, -- Можно выключить стол, не удаляя историю броней
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_restaurant_tables_restaurant_id ON restaurant_tables(restaurant_id);

-- +goose Down
DROP TABLE IF EXISTS restaurant_tables;
