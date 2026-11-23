-- +goose Up
CREATE TABLE bookings (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    table_id BIGINT NOT NULL REFERENCES restaurant_tables(id) ON DELETE CASCADE,
    guests_count INT NOT NULL CHECK (guests_count > 0),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'CONFIRMED'
        CHECK (status IN ('PENDING', 'CONFIRMED', 'CANCELLED', 'COMPLETED')),
    comment TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT check_booking_time CHECK (end_time > start_time)
);

-- Индекс для мгновенного поиска занятых столов на нужное время
CREATE INDEX idx_bookings_table_time ON bookings (table_id, start_time, end_time)
    WHERE status IN ('PENDING', 'CONFIRMED');

-- +goose Down
DROP TABLE IF EXISTS bookings;
