// Functions.go
// A complete guide to functions in Go.

package main

import (
	"fmt"
)

// 1️⃣ Simple Function (No Parameters, No Return)
func greet() {
	fmt.Println("Hello from Go functions!")
}

// 2️⃣ Function with Parameters
func addNumbers(a int, b int) {
	sum := a + b
	fmt.Println("Sum:", sum)
}

// 3️⃣ Function with Return Value
func multiply(a int, b int) int {
	return a * b
}

// 4️⃣ Multiple Return Values
func divide(a float64, b float64) (float64, string) {
	if b == 0 {
		return 0, "Error: Division by zero"
	}
	result := a / b
	return result, "Success"
}

// 5️⃣ Named Return Values
func rectangleDimensions(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

// 6️⃣ Variadic Function (Accepts multiple arguments)
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 7️⃣ Function as a Parameter
func operate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 8️⃣ Anonymous Function
func anonymousExample() {
	func(message string) {
		fmt.Println("Anonymous function says:", message)
	}("Hello, Go!")
}

// 9️⃣ Closures (Function inside a function)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 🔟 Recursive Function (Factorial)
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 1️⃣1️⃣ Defer Statement (Delays execution until surrounding function exits)
func deferExample() {
	defer fmt.Println("This is printed last (defer).")
	fmt.Println("This is printed first.")
}

// Main Function
func main() {
	// Calling simple function
	greet()

	// Calling function with parameters
	addNumbers(5, 7)

	// Calling function with return value
	product := multiply(4, 5)
	fmt.Println("Product:", product)

	// Calling function with multiple return values
	result, status := divide(10, 2)
	fmt.Println("Division Result:", result, "| Status:", status)

	// Named return values
	area, perimeter := rectangleDimensions(5, 3)
	fmt.Println("Area:", area, "| Perimeter:", perimeter)

	// Variadic function
	fmt.Println("Sum of numbers:", sumAll(1, 2, 3, 4, 5))

	// Function as a parameter
	addFunc := func(x, y int) int { return x + y }
	fmt.Println("Operation Result:", operate(10, 20, addFunc))

	// Anonymous function
	anonymousExample()

	// Closures
	nextCount := counter()
	fmt.Println("Counter:", nextCount())
	fmt.Println("Counter:", nextCount())

	// Recursive function
	fmt.Println("Factorial of 5:", factorial(5))

	// Defer statement
	deferExample()
}
