// ErrorHandling.go
// This file demonstrates Go's error handling patterns and best practices

package main

import (
	"errors"
	"fmt"
)

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

// Implement the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field %s: %s", e.Field, e.Message)
}

// Basic error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Custom error creation
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "Age cannot be negative",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "Age seems unrealistic",
		}
	}
	return nil
}

// Multiple error handling
func processUser(name string, age int) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if err := validateAge(age); err != nil {
		return fmt.Errorf("age validation failed: %w", err)
	}

	return nil
}

// Error wrapping and unwrapping
func deepFunction() error {
	return fmt.Errorf("deep error occurred")
}

func middleFunction() error {
	if err := deepFunction(); err != nil {
		return fmt.Errorf("middle function failed: %w", err)
	}
	return nil
}

func topFunction() error {
	if err := middleFunction(); err != nil {
		return fmt.Errorf("top function failed: %w", err)
	}
	return nil
}

func main() {
	// Basic error handling example
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}

	// Custom error handling
	err = validateAge(-5)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	// Multiple error handling
	err = processUser("", 25)
	if err != nil {
		fmt.Printf("Process error: %v\n", err)
	}

	// Error wrapping example
	err = topFunction()
	if err != nil {
		fmt.Printf("Error chain: %v\n", err)
	}
}
