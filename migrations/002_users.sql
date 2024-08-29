-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    first_name TEXT NOT NULL,
    second_name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    role_id INT REFERENCES roles(role_id) ON DELETE SET NULL
);

-- +migrate Down

DROP TABLE IF EXISTS users;
