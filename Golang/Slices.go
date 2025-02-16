// Slices.go
// A complete guide to slices in Go.

package main

import (
	"fmt"
)

func main() {
	// 1️⃣ Declaring a Slice
	var numbers []int
	fmt.Println("Empty slice:", numbers)

	// 2️⃣ Initializing a Slice
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println("Colors slice:", colors)

	// 3️⃣ Creating a Slice Using make()
	fruits := make([]string, 3) // length 3
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"
	fmt.Println("Fruits slice:", fruits)

	// 4️⃣ Appending to a Slice
	fruits = append(fruits, "Mango", "Peach")
	fmt.Println("Fruits after append:", fruits)

	// 5️⃣ Slicing a Slice
	subset := fruits[1:4] // from index 1 to 3
	fmt.Println("Slice subset:", subset)

	// 6️⃣ Length and Capacity
	fmt.Println("Length of fruits:", len(fruits))
	fmt.Println("Capacity of fruits:", cap(fruits))

	// 7️⃣ Copying Slices
	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)
	fmt.Println("Original slice:", original)
	fmt.Println("Copied slice:", copied)

	// 8️⃣ Iterating Over Slices (for loop)
	fmt.Print("Numbers in slice: ")
	numbers = []int{10, 20, 30, 40, 50}
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()

	// 9️⃣ Iterating with Range
	fmt.Println("Numbers using range:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 🔟 Slices of Slices (2D Slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D Slice (Matrix):", matrix)

	// 1️⃣1️⃣ Removing an Element from a Slice
	nums := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	nums = append(nums[:indexToRemove], nums[indexToRemove+1:]...)
	fmt.Println("Slice after removing element:", nums)

	// 1️⃣2️⃣ Passing Slices to Functions
	total := sumSlice(numbers)
	fmt.Println("Sum of numbers slice:", total)
}

// Function to calculate sum of a slice
func sumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}
