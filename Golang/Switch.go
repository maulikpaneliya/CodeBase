// Switch.go
// A complete guide to switch statements in Go.

package main

import (
	"fmt"
	"time"
)

func main2() {
	// 1️⃣ Simple Switch Statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's the start of the week!")
	case "Friday":
		fmt.Println("Weekend is near!")
	default:
		fmt.Println("It's a regular day.")
	}

	// 2️⃣ Multiple Cases in One Condition
	char := 'a'
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

	// 3️⃣ Switch with Expression
	num := 10
	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive number")
	}

	// 4️⃣ Switch without a Condition (like if-else ladder)
	age := 25
	switch {
	case age < 18:
		fmt.Println("Minor")
	case age >= 18 && age < 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior")
	}

	// 5️⃣ Fallthrough Keyword
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Very Good")
		fallthrough
	case "C":
		fmt.Println("Good")
	default:
		fmt.Println("Needs Improvement")
	}

	// 6️⃣ Type Switch (Checking Types of Interface Values)
	var value interface{} = 42
	switch v := value.(type) {
	case int:
		fmt.Printf("Type is int, value: %d\n", v)
	case string:
		fmt.Printf("Type is string, value: %s\n", v)
	default:
		fmt.Println("Unknown type")
	}

	// 7️⃣ Switch with Time
	currentDay := time.Now().Weekday()
	switch currentDay {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// 8️⃣ Switch with Functions
	month := time.Now().Month()
	switch month {
	case time.January, time.February, time.December:
		fmt.Println("It's winter!")
	case time.June, time.July, time.August:
		fmt.Println("It's summer!")
	default:
		fmt.Println("It's a different season.")
	}

	// 9️⃣ Nested Switch Statements
	x, y := 10, 20
	switch {
	case x > 0:
		switch {
		case y > 0:
			fmt.Println("Both x and y are positive")
		default:
			fmt.Println("x is positive, y is not")
		}
	default:
		fmt.Println("x is not positive")
	}
}
