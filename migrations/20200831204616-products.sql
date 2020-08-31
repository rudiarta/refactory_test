
-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    variant VARCHAR(255) NOT NULL,
    merchant_id INT NULL,
    price INT,
    status BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE IF EXISTS products;