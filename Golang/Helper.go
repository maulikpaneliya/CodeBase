// Helper.go
// A collection of reusable helper functions for Go projects.
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 1️⃣ Generate Random Number
func RandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// 2️⃣ Generate Random String
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(charset[rand.Intn(len(charset))])
	}
	return result.String()
}

// 3️⃣ String Contains (Case-Insensitive)
func StringContainsCaseInsensitive(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

// 4️⃣ Reverse a String
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 5️⃣ Get Current Timestamp
func CurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 6️⃣ Calculate Factorial (Recursive)
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// 7️⃣ Find Maximum Number in Slice
func MaxInSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// 8️⃣ Find Minimum Number in Slice
func MinInSlice(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

// 9️⃣ Sum of Numbers in Slice
func SumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// 🔟 Check if Number is Even
func IsEven(num int) bool {
	return num%2 == 0
}

// 1️⃣1️⃣ Check if Number is Odd
func IsOdd(num int) bool {
	return num%2 != 0
}

// 1️⃣2️⃣ Find Prime Numbers up to N
func FindPrimes(n int) []int {
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 1️⃣3️⃣ Sleep for N Seconds
func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("⏰ Slept for %d seconds\n", seconds)
}

func main() {
	// Test Helper Functions
	fmt.Println("📌 Random Number (1-100):", RandomNumber(1, 100))
	fmt.Println("📌 Random String (10 chars):", RandomString(10))
	fmt.Println("📌 'GoLang' contains 'lang'? (Case-Insensitive):", StringContainsCaseInsensitive("GoLang", "lang"))
	fmt.Println("📌 Reverse of 'hello':", ReverseString("hello"))
	fmt.Println("📌 Current Timestamp:", CurrentTimestamp())
	fmt.Println("📌 Factorial of 5:", Factorial(5))

	numbers := []int{23, 45, 67, 12, 89, 5}
	fmt.Println("📌 Max in slice:", MaxInSlice(numbers))
	fmt.Println("📌 Min in slice:", MinInSlice(numbers))
	fmt.Println("📌 Sum of slice:", SumSlice(numbers))

	fmt.Println("📌 Is 42 even?:", IsEven(42))
	fmt.Println("📌 Is 13 odd?:", IsOdd(13))

	fmt.Println("📌 Prime numbers up to 20:", FindPrimes(20))

	Sleep(2)
}
