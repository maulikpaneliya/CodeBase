// String_helper.go
// Useful String Helper Methods in Go

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 📌 1. Reverse a String
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 📌 2. Convert to Uppercase
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// 📌 3. Convert to Lowercase
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

// 📌 4. Capitalize First Letter
func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// 📌 5. Count Number of Words
func WordCount(s string) int {
	return len(strings.Fields(s))
}

// 📌 6. Check if String Contains a Substring
func ContainsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

// 📌 7. Find String Length
func StringLength(s string) int {
	return len([]rune(s)) // Properly handles Unicode
}

// 📌 8. Replace Substring
func ReplaceSubstring(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// 📌 9. Remove Whitespace
func TrimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

// 📌 10. Split String into Slices
func SplitString(s, sep string) []string {
	return strings.Split(s, sep)
}

// 📌 11. Join Slices into String
func JoinStrings(slice []string, sep string) string {
	return strings.Join(slice, sep)
}

// 📌 12. Check if String Starts with Prefix
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// 📌 13. Check if String Ends with Suffix
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// 📌 14. Count Character Occurrences
func CountCharOccurrences(s, char string) int {
	return strings.Count(s, char)
}

// 📌 15. Remove Special Characters
func RemoveSpecialCharacters(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

// 📌 16. Repeat String
func RepeatString(s string, count int) string {
	return strings.Repeat(s, count)
}

// 📌 17. Is String Empty
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// 🧪 Example Usage
func main() {
	str := "  Hello, Golang!  "
	fmt.Println("1️⃣ Reverse:", ReverseString(str))
	fmt.Println("2️⃣ Uppercase:", ToUpperCase(str))
	fmt.Println("3️⃣ Lowercase:", ToLowerCase(str))
	fmt.Println("4️⃣ Capitalize:", CapitalizeFirst("golang"))
	fmt.Println("5️⃣ Word Count:", WordCount(str))
	fmt.Println("6️⃣ Contains 'Golang':", ContainsSubstring(str, "Golang"))
	fmt.Println("7️⃣ Length:", StringLength(str))
	fmt.Println("8️⃣ Replace 'Golang' with 'Go':", ReplaceSubstring(str, "Golang", "Go"))
	fmt.Println("9️⃣ Trim Whitespace:", TrimWhitespace(str))
	fmt.Println("🔟 Split by space:", SplitString(str, " "))
	fmt.Println("🔟 Join with '-':", JoinStrings([]string{"Hello", "Golang"}, "-"))
	fmt.Println("🔟 Starts with 'Hello':", StartsWith(str, "Hello"))
	fmt.Println("🔟 Ends with '!':", EndsWith(str, "!"))
	fmt.Println("🔟 Count 'o':", CountCharOccurrences(str, "o"))
	fmt.Println("🔟 Remove Special Characters:", RemoveSpecialCharacters(str))
	fmt.Println("🔟 Repeat String:", RepeatString("Go", 3))
	fmt.Println("🔟 Is Empty:", IsEmpty("   "))
}
