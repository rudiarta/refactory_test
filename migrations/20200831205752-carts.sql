
-- +migrate Up
CREATE TABLE IF NOT EXISTS carts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    product_id INT,
    CONSTRAINT fk_user_id
    FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_product_id
    FOREIGN KEY (product_id) REFERENCES products(id)
);
-- +migrate Down
DROP TABLE IF EXISTS carts;