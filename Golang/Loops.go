// Loops.go
// A complete guide to loops in Go.

package main

import (
	"fmt"
)

func main() {
	// 1️⃣ Basic for Loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// 2️⃣ for Loop as a While Loop
	counter := 1
	for counter <= 3 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// 3️⃣ Infinite Loop
	// Uncomment with caution to avoid endless loop
	// for {
	//     fmt.Println("Infinite Loop")
	// }

	// 4️⃣ for with break Statement
	for i := 1; i <= 10; i++ {
		if i == 6 {
			fmt.Println("Loop stopped at:", i)
			break
		}
		fmt.Println(i)
	}

	// 5️⃣ for with continue Statement
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("Skipping number:", i)
			continue
		}
		fmt.Println(i)
	}

	// 6️⃣ Nested for Loops
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// 7️⃣ Looping Through Arrays
	numbers := [5]int{10, 20, 30, 40, 50}
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// 8️⃣ Looping Through Slices
	fruits := []string{"Apple", "Banana", "Cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Fruit[%d]: %s\n", index, fruit)
	}

	// 9️⃣ Looping Through Maps
	colors := map[string]string{"r": "Red", "g": "Green", "b": "Blue"}
	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 🔟 Using for with String (rune iteration)
	str := "Golang"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// 1️⃣1️⃣ Loop with Labels (to break/continue outer loops)
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("Breaking out of outer loop at i=2, j=2")
				break OuterLoop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// 1️⃣2️⃣ Sum of Numbers using a Loop
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 5:", sum)

	// 1️⃣3️⃣ Factorial using Loop
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d is %d\n", num, factorial)
}
