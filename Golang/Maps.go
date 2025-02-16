// Maps.go
// A complete guide to maps in Go.

package main

import (
	"fmt"
)

func main() {
	// 1️⃣ Creating a Map
	colors := make(map[string]string)
	colors["r"] = "Red"
	colors["g"] = "Green"
	colors["b"] = "Blue"
	fmt.Println("Colors map:", colors)

	// 2️⃣ Initializing a Map with Values
	fruits := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}
	fmt.Println("Fruits map:", fruits)

	// 3️⃣ Accessing Map Values
	fmt.Println("Color for 'g':", colors["g"])

	// 4️⃣ Updating a Value in Map
	fruits["apple"] = "green"
	fmt.Println("Updated apple color:", fruits["apple"])

	// 5️⃣ Deleting a Key from Map
	delete(colors, "b")
	fmt.Println("Colors after deletion:", colors)

	// 6️⃣ Checking if a Key Exists
	value, exists := fruits["banana"]
	if exists {
		fmt.Println("Banana color:", value)
	} else {
		fmt.Println("Banana not found")
	}

	// 7️⃣ Looping through a Map using for-range
	for key, value := range fruits {
		fmt.Printf("Fruit: %s, Color: %s\n", key, value)
	}

	// 8️⃣ Length of a Map
	fmt.Println("Number of fruits:", len(fruits))

	// 9️⃣ Map of Maps (Nested Maps)
	employees := map[string]map[string]string{
		"emp1": {
			"name": "John",
			"role": "Developer",
		},
		"emp2": {
			"name": "Alice",
			"role": "Manager",
		},
	}
	fmt.Println("Employee emp1:", employees["emp1"]["name"], "-", employees["emp1"]["role"])

	// 🔟 Map with Struct Values
	type Student struct {
		Name  string
		Grade int
	}
	students := map[string]Student{
		"101": {"Alice", 90},
		"102": {"Bob", 85},
	}
	for id, student := range students {
		fmt.Printf("ID: %s, Name: %s, Grade: %d\n", id, student.Name, student.Grade)
	}

	// 1️⃣1️⃣ Searching for a Key and Handling Missing Keys
	key := "cherry"
	if color, found := fruits[key]; found {
		fmt.Printf("Color of %s is %s\n", key, color)
	} else {
		fmt.Printf("%s not found in fruits map\n", key)
	}

	// 1️⃣2️⃣ Copying a Map
	original := map[string]int{"a": 1, "b": 2}
	duplicate := make(map[string]int)
	for k, v := range original {
		duplicate[k] = v
	}
	fmt.Println("Original map:", original)
	fmt.Println("Duplicate map:", duplicate)

	// 1️⃣3️⃣ Clearing a Map
	for k := range original {
		delete(original, k)
	}
	fmt.Println("Original map after clearing:", original)
}
