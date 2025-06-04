CREATE TABLE IF NOT EXISTS person (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);
