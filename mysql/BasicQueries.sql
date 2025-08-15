-- Basic MySQL Queries
-- This file contains commonly used MySQL queries for reference

-- Create Database
CREATE DATABASE IF NOT EXISTS example_db;
USE example_db;

-- Create Table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Basic CRUD Operations

-- Create (Insert)
INSERT INTO users (username, email) VALUES ('john_doe', 'john@example.com');
INSERT INTO users (username, email) 
VALUES 
    ('jane_doe', 'jane@example.com'),
    ('bob_smith', 'bob@example.com');

-- Read (Select)
SELECT * FROM users;
SELECT username, email FROM users WHERE id = 1;
SELECT * FROM users WHERE email LIKE '%@example.com';

-- Update
UPDATE users SET email = 'new_email@example.com' WHERE id = 1;

-- Delete
DELETE FROM users WHERE id = 1;

-- Common Clauses
SELECT * FROM users ORDER BY created_at DESC;
SELECT * FROM users LIMIT 10;
SELECT * FROM users WHERE id > 100 LIMIT 10 OFFSET 20;

-- Aggregate Functions
SELECT COUNT(*) FROM users;
SELECT MAX(id) FROM users;
SELECT username, COUNT(*) as post_count 
FROM users 
GROUP BY username 
HAVING post_count > 5;

-- Joins
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    title VARCHAR(100),
    content TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Inner Join
SELECT users.username, posts.title 
FROM users 
INNER JOIN posts ON users.id = posts.user_id;

-- Left Join
SELECT users.username, posts.title 
FROM users 
LEFT JOIN posts ON users.id = posts.user_id;

-- Subqueries
SELECT username 
FROM users 
WHERE id IN (SELECT user_id FROM posts WHERE title LIKE '%MySQL%');

-- Indexing
CREATE INDEX idx_username ON users(username);
CREATE INDEX idx_email ON users(email);

-- Views
CREATE VIEW active_users AS
SELECT * FROM users 
WHERE id IN (SELECT DISTINCT user_id FROM posts);

-- Stored Procedure
DELIMITER //
CREATE PROCEDURE GetUserPosts(IN userId INT)
BEGIN
    SELECT posts.* 
    FROM posts 
    WHERE posts.user_id = userId;
END //
DELIMITER ;

-- Call Stored Procedure
CALL GetUserPosts(1);

-- Triggers
DELIMITER //
CREATE TRIGGER before_user_delete
    BEFORE DELETE ON users
    FOR EACH ROW
BEGIN
    DELETE FROM posts WHERE user_id = OLD.id;
END //
DELIMITER ;
