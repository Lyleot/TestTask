-- +migrate Up

CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- +migrate Down

DROP TABLE IF EXISTS roles;