// Array.go
// A complete guide to arrays in Go.

package main

import (
	"fmt"
)

func main() {
	// 1️⃣ Declaring an Array
	var numbers [5]int
	fmt.Println("Empty array:", numbers)

	// 2️⃣ Initializing an Array
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println("Colors array:", colors)

	// 3️⃣ Short-hand Initialization
	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println("Primes array:", primes)

	// 4️⃣ Accessing Elements
	fmt.Println("First color:", colors[0])
	fmt.Println("Second prime:", primes[1])

	// 5️⃣ Modifying Elements
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println("Modified array:", numbers)

	// 6️⃣ Array Length
	fmt.Println("Length of primes array:", len(primes))

	// 7️⃣ Iterating Over Arrays (for loop)
	fmt.Print("Numbers array: ")
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()

	// 8️⃣ Iterating with Range
	fmt.Println("Colors array using range:")
	for index, value := range colors {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// 9️⃣ Multidimensional Arrays
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[1][2] = 5
	fmt.Println("2D Array (Matrix):", matrix)

	// Initializing a 2D Array
	grid := [2][2]string{
		{"X", "O"},
		{"O", "X"},
	}
	fmt.Println("Grid:", grid)

	// 🔟 Array Copying (Value Type)
	original := [3]int{1, 2, 3}
	copyArray := original
	copyArray[0] = 100
	fmt.Println("Original Array:", original)
	fmt.Println("Copied Array:", copyArray)

	// 1️⃣1️⃣ Array Comparison
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println("Are arrays equal?", arr1 == arr2)

	// 1️⃣2️⃣ Array with Functions
	sum := sumArray(primes)
	fmt.Println("Sum of primes:", sum)
}

// Function to calculate sum of an integer array
func sumArray(arr [5]int) int {
	total := 0
	for _, value := range arr {
		total += value
	}
	return total
}
