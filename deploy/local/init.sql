
CREATE DATABASE userdb;
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       age INTEGER NOT NULL CHECK (age > 0)
);

CREATE INDEX idx_users_email ON users(email);

CREATE DATABASE orderdb;

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(200) NOT NULL,
                        description TEXT NOT NULL,
                        user_id INTEGER NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_orders_user_id ON orders(user_id);