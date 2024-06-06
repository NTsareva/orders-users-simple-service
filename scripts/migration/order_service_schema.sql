\c orderdb;

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(200) NOT NULL,
                        description TEXT NOT NULL,
                        user_id INTEGER NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_orders_id ON orders(id);
CREATE INDEX idx_orders_user_id ON orders(user_id);