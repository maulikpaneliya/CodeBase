-- MySQL CRUD Operations

-- 1. INSERT Operations
-- ------------------

-- Simple INSERT
INSERT INTO users (username, email, password_hash, first_name, last_name)
VALUES ('johndoe', 'john@example.com', 'hashed_password', 'John', 'Doe');

-- Multiple row INSERT
INSERT INTO products (name, description, price, stock_quantity, category_id)
VALUES 
    ('Laptop', 'High-performance laptop', 999.99, 50, 1),
    ('Mouse', 'Wireless mouse', 29.99, 100, 1),
    ('Keyboard', 'Mechanical keyboard', 89.99, 75, 1);

-- INSERT with SELECT
INSERT INTO backup_users
SELECT * FROM users
WHERE created_at < DATE_SUB(NOW(), INTERVAL 1 YEAR);

-- 2. SELECT Operations
-- ------------------

-- Basic SELECT
SELECT id, username, email, created_at
FROM users
WHERE is_active = true
ORDER BY created_at DESC;

-- SELECT with JOIN
SELECT 
    o.id AS order_id,
    u.username,
    p.name AS product_name,
    oi.quantity,
    oi.unit_price,
    (oi.quantity * oi.unit_price) AS total_price
FROM orders o
JOIN users u ON o.user_id = u.id
JOIN order_items oi ON o.id = oi.order_id
JOIN products p ON oi.product_id = p.id
WHERE o.status = 'shipped'
ORDER BY o.order_date DESC;

-- Aggregation
SELECT 
    category_id,
    COUNT(*) AS total_products,
    AVG(price) AS avg_price,
    MIN(price) AS min_price,
    MAX(price) AS max_price,
    SUM(stock_quantity) AS total_stock
FROM products
GROUP BY category_id
HAVING total_products > 5;

-- Subquery
SELECT *
FROM products
WHERE price > (
    SELECT AVG(price)
    FROM products
    WHERE category_id = 1
);

-- 3. UPDATE Operations
-- ------------------

-- Simple UPDATE
UPDATE products
SET price = price * 1.1
WHERE category_id = 1;

-- UPDATE with JOIN
UPDATE orders o
JOIN order_items oi ON o.id = oi.order_id
SET o.total_amount = (
    SELECT SUM(quantity * unit_price)
    FROM order_items
    WHERE order_id = o.id
)
WHERE o.status = 'pending';

-- UPDATE with subquery
UPDATE users
SET is_active = false
WHERE id IN (
    SELECT user_id
    FROM orders
    GROUP BY user_id
    HAVING COUNT(*) = 0
);

-- 4. DELETE Operations
-- ------------------

-- Simple DELETE
DELETE FROM products
WHERE stock_quantity = 0 AND is_active = false;

-- DELETE with JOIN
DELETE o, oi
FROM orders o
JOIN order_items oi ON o.id = oi.order_id
WHERE o.status = 'cancelled'
AND o.order_date < DATE_SUB(NOW(), INTERVAL 1 MONTH);

-- DELETE with subquery
DELETE FROM users
WHERE id IN (
    SELECT user_id
    FROM orders
    GROUP BY user_id
    HAVING COUNT(*) = 0
);

-- 5. Complex Queries
-- ----------------

-- Window Functions
SELECT 
    category_id,
    name,
    price,
    AVG(price) OVER (PARTITION BY category_id) AS avg_category_price,
    price - AVG(price) OVER (PARTITION BY category_id) AS price_diff,
    RANK() OVER (PARTITION BY category_id ORDER BY price DESC) AS price_rank
FROM products;

-- Common Table Expression (CTE)
WITH monthly_sales AS (
    SELECT 
        DATE_FORMAT(order_date, '%Y-%m') AS month,
        SUM(total_amount) AS total_sales
    FROM orders
    WHERE status = 'delivered'
    GROUP BY DATE_FORMAT(order_date, '%Y-%m')
)
SELECT 
    month,
    total_sales,
    LAG(total_sales) OVER (ORDER BY month) AS prev_month_sales,
    ((total_sales - LAG(total_sales) OVER (ORDER BY month)) / 
     LAG(total_sales) OVER (ORDER BY month) * 100) AS growth_percentage
FROM monthly_sales
ORDER BY month;

-- Best Practices:
-- 1. Always use WHERE clauses to limit data
-- 2. Use appropriate indexes for queries
-- 3. Use transactions for multiple operations
-- 4. Be careful with DELETE operations
-- 5. Use JOIN instead of subqueries when possible
-- 6. Use meaningful column aliases
-- 7. Group related queries together
-- 8. Use appropriate ORDER BY for sorted results
-- 9. Use LIMIT for large result sets
-- 10. Consider query performance impact
