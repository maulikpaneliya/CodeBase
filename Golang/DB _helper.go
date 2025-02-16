// DB_Helper.go
// Common Database Helper Methods for CRUD Operations in Go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 1️⃣ Create Database Connection
func ConnectDB() (*sql.DB, error) {
	// Replace with your DB credentials
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("✅ Database Connected")
	return db, nil
}

// 2️⃣ Common Select Method
func SelectAll(db *sql.DB, tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	return db.Query(query)
}

// 3️⃣ Common Insert Method
func InsertRecord(db *sql.DB, tableName string, data interface{}) error {
	columns, values, placeholders := getFieldsAndValues(data)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName, strings.Join(columns, ","), strings.Join(placeholders, ","))
	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}
	fmt.Println("✅ Record Inserted")
	return nil
}

// 4️⃣ Common Update Method
func UpdateRecord(db *sql.DB, tableName string, data interface{}, primaryKey string, primaryValue interface{}) error {
	columns, values, _ := getFieldsAndValues(data)
	var setClauses []string
	for _, col := range columns {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?",
		tableName, strings.Join(setClauses, ","), primaryKey)
	values = append(values, primaryValue)

	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}
	fmt.Println("✅ Record Updated")
	return nil
}

// 5️⃣ Common Delete Method
func DeleteRecord(db *sql.DB, tableName, primaryKey string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableName, primaryKey)
	_, err := db.Exec(query, value)
	if err != nil {
		return err
	}
	fmt.Println("✅ Record Deleted")
	return nil
}

// 6️⃣ Check If Record Exists
func CheckIfExists(db *sql.DB, tableName, columnName string, value interface{}) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE %s = ? LIMIT 1", tableName, columnName)
	var exists bool
	err := db.QueryRow(query, value).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// Helper Function: Get Struct Fields and Values
func getFieldsAndValues(data interface{}) ([]string, []interface{}, []string) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	var columns []string
	var values []interface{}
	var placeholders []string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		column := field.Tag.Get("db")
		if column == "" {
			continue
		}
		columns = append(columns, column)
		values = append(values, v.Field(i).Interface())
		placeholders = append(placeholders, "?")
	}

	return columns, values, placeholders
}

// Example Usage
type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("DB Connection Failed:", err)
	}
	defer db.Close()

	// Insert Example
	user := User{Name: "John", Email: "john@example.com"}
	InsertRecord(db, "users", user)

	// Update Example
	updatedUser := User{Name: "John Updated", Email: "john.updated@example.com"}
	UpdateRecord(db, "users", updatedUser, "id", 1)

	// Check If Exists
	exists, _ := CheckIfExists(db, "users", "email", "john.updated@example.com")
	fmt.Println("User exists:", exists)

	// Delete Example
	DeleteRecord(db, "users", "id", 1)
}
