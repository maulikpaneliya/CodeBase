-- MySQL Joins and Relationships

-- 1. INNER JOIN
-- ------------

-- Basic INNER JOIN
SELECT 
    p.id,
    p.name AS product_name,
    c.name AS category_name,
    p.price
FROM products p
INNER JOIN categories c ON p.category_id = c.id;

-- Multiple INNER JOINs
SELECT 
    o.id AS order_id,
    u.username,
    p.name AS product_name,
    oi.quantity,
    oi.unit_price
FROM orders o
INNER JOIN users u ON o.user_id = u.id
INNER JOIN order_items oi ON o.id = oi.order_id
INNER JOIN products p ON oi.product_id = p.id;

-- 2. LEFT JOIN
-- -----------

-- Basic LEFT JOIN
SELECT 
    c.name AS category_name,
    COUNT(p.id) AS product_count
FROM categories c
LEFT JOIN products p ON c.id = p.category_id
GROUP BY c.id, c.name;

-- LEFT JOIN with multiple conditions
SELECT 
    u.username,
    COUNT(o.id) AS total_orders,
    COALESCE(SUM(o.total_amount), 0) AS total_spent
FROM users u
LEFT JOIN orders o ON u.id = o.user_id AND o.status = 'delivered'
GROUP BY u.id, u.username;

-- 3. RIGHT JOIN
-- ------------

-- Basic RIGHT JOIN
SELECT 
    p.name AS product_name,
    COALESCE(SUM(oi.quantity), 0) AS total_ordered
FROM order_items oi
RIGHT JOIN products p ON oi.product_id = p.id
GROUP BY p.id, p.name;

-- 4. FULL JOIN (Emulated in MySQL)
-- ------------------------------

-- Emulating FULL JOIN using UNION
SELECT c.name AS category_name, p.name AS product_name
FROM categories c
LEFT JOIN products p ON c.id = p.category_id
UNION
SELECT c.name AS category_name, p.name AS product_name
FROM categories c
RIGHT JOIN products p ON c.id = p.category_id;

-- 5. CROSS JOIN
-- ------------

-- Generate price matrix
SELECT 
    p1.name AS product_1,
    p2.name AS product_2,
    (p1.price + p2.price) AS combo_price
FROM products p1
CROSS JOIN products p2
WHERE p1.id < p2.id;

-- 6. SELF JOIN
-- -----------

-- Employee hierarchy
SELECT 
    e1.name AS employee,
    e2.name AS manager
FROM employees e1
LEFT JOIN employees e2 ON e1.manager_id = e2.id;

-- Product recommendations
SELECT 
    p1.name AS product,
    p2.name AS related_product,
    p2.price
FROM products p1
INNER JOIN products p2 ON p1.category_id = p2.category_id
    AND p1.id != p2.id
WHERE p1.id = 1;

-- 7. Complex JOIN Patterns
-- ----------------------

-- Order summary with multiple aggregations
SELECT 
    u.username,
    COUNT(DISTINCT o.id) AS total_orders,
    SUM(oi.quantity) AS total_items,
    SUM(oi.quantity * oi.unit_price) AS total_spent,
    MAX(o.order_date) AS last_order_date
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
LEFT JOIN order_items oi ON o.id = oi.order_id
GROUP BY u.id, u.username;

-- Category performance analysis
SELECT 
    c.name AS category_name,
    COUNT(DISTINCT p.id) AS total_products,
    COUNT(DISTINCT o.id) AS orders_with_category,
    SUM(oi.quantity) AS total_items_sold,
    SUM(oi.quantity * oi.unit_price) AS total_revenue,
    AVG(p.price) AS avg_price
FROM categories c
LEFT JOIN products p ON c.id = p.category_id
LEFT JOIN order_items oi ON p.id = oi.product_id
LEFT JOIN orders o ON oi.order_id = o.id
WHERE o.status = 'delivered' OR o.status IS NULL
GROUP BY c.id, c.name;

-- Customer purchase patterns
SELECT 
    u.username,
    c.name AS category_name,
    COUNT(DISTINCT o.id) AS orders_in_category,
    SUM(oi.quantity) AS items_bought,
    AVG(oi.unit_price) AS avg_item_price
FROM users u
CROSS JOIN categories c
LEFT JOIN orders o ON u.id = o.user_id
LEFT JOIN order_items oi ON o.id = oi.order_id
LEFT JOIN products p ON oi.product_id = p.id AND p.category_id = c.id
GROUP BY u.id, u.username, c.id, c.name;

-- Best Practices:
-- 1. Always specify JOIN type explicitly
-- 2. Use appropriate indexes on JOIN columns
-- 3. Use table aliases for better readability
-- 4. Consider JOIN order for performance
-- 5. Use WHERE clauses after JOINs
-- 6. Include all necessary columns in GROUP BY
-- 7. Use meaningful column aliases
-- 8. Consider using views for complex joins
-- 9. Test JOIN performance with EXPLAIN
-- 10. Use appropriate data types for JOIN columns
