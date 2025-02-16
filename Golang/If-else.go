// If-else.go
// A complete guide to if-else statements in Go.

package main

import (
	"fmt"
)

func main() {
	// 1️⃣ Simple if Statement
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5")
	}

	// 2️⃣ if-else Statement
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// 3️⃣ if-else if-else Statement
	score := 75
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// 4️⃣ if with Short Statement
	if length := len("Golang"); length > 5 {
		fmt.Println("Long string")
	} else {
		fmt.Println("Short string")
	}

	// 5️⃣ Nested if Statements
	marks := 85
	if marks >= 50 {
		fmt.Println("Passed")
		if marks >= 80 {
			fmt.Println("Excellent Score!")
		}
	} else {
		fmt.Println("Failed")
	}

	// 6️⃣ Using Logical Operators
	x, y := 15, 20
	if x > 10 && y > 15 {
		fmt.Println("Both conditions are true")
	}

	if x > 10 || y < 15 {
		fmt.Println("At least one condition is true")
	}

	// 7️⃣ Comparing Strings
	city := "Chicago"
	if city == "Chicago" {
		fmt.Println("Welcome to Chicago!")
	} else {
		fmt.Println("You're in another city")
	}

	// 8️⃣ Check Even or Odd
	number := 7
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 9️⃣ Check Positive, Negative, or Zero
	n := -5
	if n > 0 {
		fmt.Println("Positive number")
	} else if n < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("It's zero")
	}

	// 🔟 Using if with a Function
	temperature := 30
	if isHot(temperature) {
		fmt.Println("It's a hot day")
	} else {
		fmt.Println("The weather is mild")
	}
}

// Function to check if temperature is hot
func isHot(temp int) bool {
	return temp >= 30
}
