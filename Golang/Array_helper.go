// Array_helper.go
// Useful Array, Slice, and Map Helper Methods in Go

package main

import (
	"fmt"
	"sort"
)

// 📌 1. Find Maximum Value in an Array
func FindMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

// 📌 2. Find Minimum Value in an Array
func FindMin(arr []int) int {
	min := arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

// 📌 3. Check if Element Exists in Slice
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// 📌 4. Reverse a Slice
func ReverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))
	for i, v := range slice {
		reversed[len(slice)-1-i] = v
	}
	return reversed
}

// 📌 5. Remove Duplicates from Slice
func RemoveDuplicates(slice []int) []int {
	unique := make(map[int]bool)
	var result []int
	for _, v := range slice {
		if !unique[v] {
			unique[v] = true
			result = append(result, v)
		}
	}
	return result
}

// 📌 6. Sort a Slice in Ascending Order
func SortAscending(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	return sorted
}

// 📌 7. Sort a Slice in Descending Order
func SortDescending(arr []int) []int {
	sorted := SortAscending(arr)
	// Reverse to get descending order
	return ReverseSlice(sorted)
}

// 📌 8. Merge Two Slices
func MergeSlices(slice1, slice2 []int) []int {
	return append(slice1, slice2...)
}

// 📌 9. Sum of Elements in Slice
func SumOfSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// 📌 10. Get Keys from Map
func GetMapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 📌 11. Get Values from Map
func GetMapValues(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// 📌 12. Merge Two Maps
func MergeMaps(map1, map2 map[string]int) map[string]int {
	merged := make(map[string]int)
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}

// 📌 13. Check if Key Exists in Map
func KeyExists(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

// 📌 14. Find Length of Slice or Map
func GetLength(data interface{}) int {
	switch v := data.(type) {
	case []int:
		return len(v)
	case []string:
		return len(v)
	case map[string]int:
		return len(v)
	default:
		return -1
	}
}

// 🧪 Example Usage
func main() {
	arr := []int{10, 20, 20, 30, 40, 10}
	strSlice := []string{"apple", "banana", "cherry"}
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}

	fmt.Println("1️⃣ Max Value:", FindMax(arr))
	fmt.Println("2️⃣ Min Value:", FindMin(arr))
	fmt.Println("3️⃣ Contains 'banana':", Contains(strSlice, "banana"))
	fmt.Println("4️⃣ Reverse Slice:", ReverseSlice(arr))
	fmt.Println("5️⃣ Remove Duplicates:", RemoveDuplicates(arr))
	fmt.Println("6️⃣ Sort Ascending:", SortAscending(arr))
	fmt.Println("7️⃣ Sort Descending:", SortDescending(arr))
	fmt.Println("8️⃣ Merge Slices:", MergeSlices(arr, []int{50, 60}))
	fmt.Println("9️⃣ Sum of Slice:", SumOfSlice(arr))
	fmt.Println("🔟 Map Keys:", GetMapKeys(map1))
	fmt.Println("🔟 Map Values:", GetMapValues(map1))
	fmt.Println("🔟 Merged Maps:", MergeMaps(map1, map2))
	fmt.Println("🔟 Key 'a' Exists:", KeyExists(map1, "a"))
	fmt.Println("🔟 Length of arr:", GetLength(arr))
	fmt.Println("🔟 Length of map1:", GetLength(map1))
}
