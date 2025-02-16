// TypeConversion.go
// A complete guide to type conversions in Go.

package main

import (
	"fmt"
	"strconv"
)

// 1️⃣ Integer to Float
func intToFloat() {
	var intVal int = 42
	var floatVal float64 = float64(intVal)
	fmt.Println("Integer to Float:", floatVal)
}

// 2️⃣ Float to Integer (Truncates decimal part)
func floatToInt() {
	var floatVal float64 = 42.78
	var intVal int = int(floatVal)
	fmt.Println("Float to Integer (Truncated):", intVal)
}

// 3️⃣ Integer to String
func intToString() {
	var intVal int = 123
	var strVal string = strconv.Itoa(intVal) // Using strconv package
	fmt.Println("Integer to String:", strVal)
}

// 4️⃣ String to Integer
func stringToInt() {
	var strVal string = "123"
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("String to Integer:", intVal)
	}
}

// 5️⃣ Float to String
func floatToString() {
	var floatVal float64 = 45.67
	var strVal string = fmt.Sprintf("%f", floatVal) // Using fmt package
	fmt.Println("Float to String:", strVal)
}

// 6️⃣ String to Float
func stringToFloat() {
	var strVal string = "45.67"
	floatVal, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("String to Float:", floatVal)
	}
}

// 7️⃣ Boolean to String
func boolToString() {
	var boolVal bool = true
	var strVal string = strconv.FormatBool(boolVal)
	fmt.Println("Boolean to String:", strVal)
}

// 8️⃣ String to Boolean
func stringToBool() {
	var strVal string = "true"
	boolVal, err := strconv.ParseBool(strVal)
	if err != nil {
		fmt.Println("Error converting string to boolean:", err)
	} else {
		fmt.Println("String to Boolean:", boolVal)
	}
}

// 9️⃣ Byte to String
func byteToString() {
	var byteArr = []byte{'H', 'e', 'l', 'l', 'o'}
	var strVal string = string(byteArr)
	fmt.Println("Byte to String:", strVal)
}

// 🔟 String to Byte
func stringToByte() {
	var strVal string = "Hello"
	var byteArr []byte = []byte(strVal)
	fmt.Println("String to Byte:", byteArr)
}

// 1️⃣1️⃣ Rune to String
func runeToString() {
	var runeVal rune = 'A'
	var strVal string = string(runeVal)
	fmt.Println("Rune to String:", strVal)
}

// 1️⃣2️⃣ String to Rune
func stringToRune() {
	var strVal string = "Hello"
	for i, r := range strVal {
		fmt.Printf("Character at %d: %c (rune)\n", i, r)
	}
}

// 1️⃣3️⃣ Type Assertion (Interface to Concrete Type)
func typeAssertionExample() {
	var i interface{} = "Hello, Go"
	str, ok := i.(string)
	if ok {
		fmt.Println("Type Assertion Successful:", str)
	} else {
		fmt.Println("Type Assertion Failed")
	}
}

// 1️⃣4️⃣ Type Conversion with Pointers
func pointerTypeConversion() {
	var x int = 10
	var ptr *int = &x
	fmt.Println("Pointer Value:", *ptr)
	fmt.Println("Pointer Address:", ptr)
}

func main() {
	fmt.Println("\n🔹 Integer to Float:")
	intToFloat()

	fmt.Println("\n🔹 Float to Integer:")
	floatToInt()

	fmt.Println("\n🔹 Integer to String:")
	intToString()

	fmt.Println("\n🔹 String to Integer:")
	stringToInt()

	fmt.Println("\n🔹 Float to String:")
	floatToString()

	fmt.Println("\n🔹 String to Float:")
	stringToFloat()

	fmt.Println("\n🔹 Boolean to String:")
	boolToString()

	fmt.Println("\n🔹 String to Boolean:")
	stringToBool()

	fmt.Println("\n🔹 Byte to String:")
	byteToString()

	fmt.Println("\n🔹 String to Byte:")
	stringToByte()

	fmt.Println("\n🔹 Rune to String:")
	runeToString()

	fmt.Println("\n🔹 String to Rune:")
	stringToRune()

	fmt.Println("\n🔹 Type Assertion (Interface to Concrete Type):")
	typeAssertionExample()

	fmt.Println("\n🔹 Type Conversion with Pointers:")
	pointerTypeConversion()
}
