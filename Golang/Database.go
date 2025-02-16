// Database.go
// A complete guide to database operations in Go with MySQL and SQLite.

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/mattn/go-sqlite3"    // SQLite driver
)

// 1️⃣ MySQL Connection
func connectMySQL() (*sql.DB, error) {
	// Change username, password, and dbname as per your setup
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("✅ Connected to MySQL database")
	return db, nil
}

// 2️⃣ SQLite Connection
func connectSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./testdb.sqlite")
	if err != nil {
		return nil, err
	}
	fmt.Println("✅ Connected to SQLite database")
	return db, nil
}

// 3️⃣ Create Table
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER,
		email TEXT UNIQUE
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("✅ Users table created successfully")
}

// 4️⃣ Insert Record
func insertUser(db *sql.DB, name string, age int, email string) {
	query := `INSERT INTO users (name, age, email) VALUES (?, ?, ?)`
	_, err := db.Exec(query, name, age, email)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	fmt.Println("✅ User added successfully:", name)
}

// 5️⃣ Read Records
func readUsers(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name, age, email FROM users`)
	if err != nil {
		log.Fatalf("Error reading users: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n📜 User List:")
	for rows.Next() {
		var id, age int
		var name, email string
		if err := rows.Scan(&id, &name, &age, &email); err != nil {
			log.Fatalf("Error scanning user: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s\n", id, name, age, email)
	}
}

// 6️⃣ Update Record
func updateUser(db *sql.DB, id int, newEmail string) {
	query := `UPDATE users SET email = ? WHERE id = ?`
	_, err := db.Exec(query, newEmail, id)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Println("✅ User email updated successfully")
}

// 7️⃣ Delete Record
func deleteUser(db *sql.DB, id int) {
	query := `DELETE FROM users WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	fmt.Println("✅ User deleted successfully")
}

func main() {
	// Choose Database: MySQL or SQLite
	db, err := connectSQLite() // Use connectMySQL() for MySQL
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Database Operations
	createTable(db)
	insertUser(db, "Alice", 25, "alice@example.com")
	insertUser(db, "Bob", 30, "bob@example.com")

	readUsers(db)

	updateUser(db, 1, "alice.new@example.com")
	readUsers(db)

	deleteUser(db, 2)
	readUsers(db)
}
