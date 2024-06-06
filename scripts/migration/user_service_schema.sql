\c userdb;

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       age INTEGER NOT NULL CHECK (age > 0)
);

CREATE UNIQUE INDEX idx_users_id ON users(id);
CREATE UNIQUE INDEX idx_users_email ON users(email);