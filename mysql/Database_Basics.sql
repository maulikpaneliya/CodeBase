-- MySQL Basic Database Operations

-- 1. Database Operations
-- --------------------

-- Create a new database
CREATE DATABASE IF NOT EXISTS ecommerce_db
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

-- Switch to the database
USE ecommerce_db;

-- 2. Table Creation
-- ---------------

-- Create Users table
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_username (username)
) ENGINE=InnoDB;

-- Create Products table
CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INT NOT NULL DEFAULT 0,
    category_id INT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
    INDEX idx_category (category_id),
    INDEX idx_name (name)
) ENGINE=InnoDB;

-- Create Orders table
CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    status ENUM('pending', 'processing', 'shipped', 'delivered', 'cancelled') DEFAULT 'pending',
    total_amount DECIMAL(10, 2) NOT NULL,
    shipping_address TEXT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user (user_id),
    INDEX idx_status (status)
) ENGINE=InnoDB;

-- Create Order Items table
CREATE TABLE order_items (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE RESTRICT,
    UNIQUE KEY unique_order_product (order_id, product_id)
) ENGINE=InnoDB;

-- 3. Table Modifications
-- --------------------

-- Add a column
ALTER TABLE users
ADD COLUMN phone_number VARCHAR(20);

-- Modify a column
ALTER TABLE products
MODIFY COLUMN price DECIMAL(12, 2);

-- Add an index
ALTER TABLE orders
ADD INDEX idx_order_date (order_date);

-- 4. Constraints
-- ------------

-- Add a CHECK constraint
ALTER TABLE products
ADD CONSTRAINT chk_price CHECK (price >= 0);

-- Add a FOREIGN KEY constraint
ALTER TABLE order_items
ADD CONSTRAINT fk_order_items_product
FOREIGN KEY (product_id) REFERENCES products(id)
ON DELETE RESTRICT
ON UPDATE CASCADE;

-- 5. Views
-- -------

-- Create a view for order summaries
CREATE VIEW order_summaries AS
SELECT 
    o.id AS order_id,
    u.username,
    o.total_amount,
    o.status,
    o.order_date,
    COUNT(oi.id) AS total_items
FROM orders o
JOIN users u ON o.user_id = u.id
JOIN order_items oi ON o.id = oi.order_id
GROUP BY o.id;

-- 6. Stored Procedure
-- -----------------

DELIMITER //

CREATE PROCEDURE create_order(
    IN p_user_id INT,
    IN p_shipping_address TEXT,
    OUT p_order_id INT
)
BEGIN
    DECLARE exit handler FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Error creating order';
    END;

    START TRANSACTION;
    
    INSERT INTO orders (user_id, shipping_address, total_amount)
    VALUES (p_user_id, p_shipping_address, 0);
    
    SET p_order_id = LAST_INSERT_ID();
    
    COMMIT;
END //

DELIMITER ;

-- Best Practices:
-- 1. Always specify character set and collation
-- 2. Use appropriate data types and lengths
-- 3. Include proper indexes for frequently queried columns
-- 4. Implement foreign key constraints for data integrity
-- 5. Use ENUM for fixed sets of values
-- 6. Include created_at and updated_at timestamps
-- 7. Use meaningful names for tables, columns, and constraints
-- 8. Add appropriate indexes for foreign keys
-- 9. Use transactions for data consistency
-- 10. Document complex procedures and views
