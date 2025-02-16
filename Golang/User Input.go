// UserInput.go
// A complete guide to handling user input in Go.

package main

import (
	"bufio"   // For buffered input
	"fmt"     // For formatted I/O
	"os"      // For operating system functionalities
	"strconv" // For string to number conversions
	"strings" // For string operations
)

func main() {
	// 1️⃣ Simple User Input (fmt.Scanln)
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // Reads input until the first space or newline
	fmt.Println("Hello,", name)

	// 2️⃣ Multiple Inputs (fmt.Scan)
	var age int
	var city string
	fmt.Print("Enter your age and city: ")
	fmt.Scan(&age, &city) // Reads multiple values separated by space
	fmt.Println("You are", age, "years old from", city)

	// 3️⃣ Input with Spaces (bufio.Reader)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full address: ")
	address, _ := reader.ReadString('\n') // Reads until newline
	address = strings.TrimSpace(address)  // Removes trailing newline
	fmt.Println("Your address:", address)

	// 4️⃣ String to Integer Conversion
	fmt.Print("Enter your lucky number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Removes newline character

	// Convert string to integer
	luckyNumber, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number. Please enter a valid integer.")
	} else {
		fmt.Println("Your lucky number is:", luckyNumber)
	}

	// 5️⃣ String to Float Conversion
	fmt.Print("Enter your height (in meters): ")
	heightInput, _ := reader.ReadString('\n')
	heightInput = strings.TrimSpace(heightInput)

	height, err := strconv.ParseFloat(heightInput, 64)
	if err != nil {
		fmt.Println("Invalid height. Please enter a valid number.")
	} else {
		fmt.Printf("Your height is: %.2f meters\n", height)
	}

	// 6️⃣ Boolean Input
	fmt.Print("Are you a student? (true/false): ")
	studentInput, _ := reader.ReadString('\n')
	studentInput = strings.TrimSpace(studentInput)

	isStudent, err := strconv.ParseBool(studentInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter true or false.")
	} else {
		if isStudent {
			fmt.Println("Great! Keep learning.")
		} else {
			fmt.Println("Lifelong learning is important!")
		}
	}
}
