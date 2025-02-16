// String.go
// A complete guide to strings in Go.

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 1️⃣ Basic String Declaration
func basicStringExample() {
	var str1 string = "Hello, Go!"
	str2 := "Welcome to Golang"
	fmt.Println("String 1:", str1)
	fmt.Println("String 2:", str2)
}

// 2️⃣ String Length
func stringLengthExample() {
	str := "Hello, 世界"
	fmt.Println("Byte Length:", len(str))                         // Byte length
	fmt.Println("Character Length:", utf8.RuneCountInString(str)) // Unicode length
}

// 3️⃣ String Concatenation
func stringConcatExample() {
	first := "Golang"
	second := "Tutorial"
	combined := first + " " + second
	fmt.Println("Concatenated String:", combined)
}

// 4️⃣ String Formatting
func stringFormattingExample() {
	name := "John"
	age := 30
	formatted := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println("Formatted String:", formatted)
}

// 5️⃣ String Comparison
func stringComparisonExample() {
	str1 := "Hello"
	str2 := "hello"
	fmt.Println("Case-sensitive Comparison:", str1 == str2)
	fmt.Println("Case-insensitive Comparison:", strings.EqualFold(str1, str2))
}

// 6️⃣ Substring Check
func substringExample() {
	str := "Learn Golang with examples"
	substr := "Golang"
	fmt.Println("Contains substring:", strings.Contains(str, substr))
	fmt.Println("Contains any character:", strings.ContainsAny(str, "xyz"))
	fmt.Println("Has prefix:", strings.HasPrefix(str, "Learn"))
	fmt.Println("Has suffix:", strings.HasSuffix(str, "examples"))
}

// 7️⃣ String Manipulations
func stringManipulationExample() {
	str := "  Hello, Golang!  "
	fmt.Println("Original:", str)
	fmt.Println("To Upper:", strings.ToUpper(str))
	fmt.Println("To Lower:", strings.ToLower(str))
	fmt.Println("Trim Spaces:", strings.TrimSpace(str))
	fmt.Println("Replace All:", strings.ReplaceAll(str, "Golang", "Go"))
}

// 8️⃣ Splitting Strings
func stringSplitExample() {
	str := "apple,banana,orange"
	parts := strings.Split(str, ",")
	fmt.Println("Split Result:", parts)
	for i, part := range parts {
		fmt.Printf("Part %d: %s\n", i+1, part)
	}
}

// 9️⃣ Joining Strings
func stringJoinExample() {
	words := []string{"Go", "is", "awesome"}
	joined := strings.Join(words, " ")
	fmt.Println("Joined String:", joined)
}

// 🔟 String Iteration
func stringIterationExample() {
	str := "Golang"
	fmt.Println("Iterating over string:")
	for i, ch := range str {
		fmt.Printf("Index: %d, Character: %c\n", i, ch)
	}
}

// 1️⃣1️⃣ String Runes (Unicode Handling)
func stringRuneExample() {
	str := "Hello, 世界"
	fmt.Println("Rune iteration:")
	for i, r := range str {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}
}

// 1️⃣2️⃣ String Conversion
func stringConversionExample() {
	// Int to String
	num := 123
	str := fmt.Sprintf("%d", num)
	fmt.Println("Int to String:", str)

	// String to Byte Array
	bytes := []byte("Hello")
	fmt.Println("String to Bytes:", bytes)

	// Byte Array to String
	strFromBytes := string(bytes)
	fmt.Println("Bytes to String:", strFromBytes)
}

// 1️⃣3️⃣ String Searching
func stringSearchExample() {
	str := "Golang is powerful"
	fmt.Println("Index of 'powerful':", strings.Index(str, "powerful"))
	fmt.Println("Index of 'Go':", strings.Index(str, "Go"))
	fmt.Println("Last Index of 'o':", strings.LastIndex(str, "o"))
}

// 1️⃣4️⃣ String Repetition
func stringRepeatExample() {
	str := "Go! "
	repeated := strings.Repeat(str, 3)
	fmt.Println("Repeated String:", repeated)
}

// 1️⃣5️⃣ Multiline Strings
func multilineStringExample() {
	multiStr := `This is a
multiline
string in Go.`
	fmt.Println("Multiline String:")
	fmt.Println(multiStr)
}

func main() {
	fmt.Println("\n🔹 Basic String Declaration:")
	basicStringExample()

	fmt.Println("\n🔹 String Length:")
	stringLengthExample()

	fmt.Println("\n🔹 String Concatenation:")
	stringConcatExample()

	fmt.Println("\n🔹 String Formatting:")
	stringFormattingExample()

	fmt.Println("\n🔹 String Comparison:")
	stringComparisonExample()

	fmt.Println("\n🔹 Substring Check:")
	substringExample()

	fmt.Println("\n🔹 String Manipulations:")
	stringManipulationExample()

	fmt.Println("\n🔹 Splitting Strings:")
	stringSplitExample()

	fmt.Println("\n🔹 Joining Strings:")
	stringJoinExample()

	fmt.Println("\n🔹 String Iteration:")
	stringIterationExample()

	fmt.Println("\n🔹 String Runes (Unicode Handling):")
	stringRuneExample()

	fmt.Println("\n🔹 String Conversion:")
	stringConversionExample()

	fmt.Println("\n🔹 String Searching:")
	stringSearchExample()

	fmt.Println("\n🔹 String Repetition:")
	stringRepeatExample()

	fmt.Println("\n🔹 Multiline Strings:")
	multilineStringExample()
}
