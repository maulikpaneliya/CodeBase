// JsonWithGo.go
// A comprehensive guide to working with JSON in Go.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 1️⃣ Struct for JSON Serialization and Deserialization
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

// 2️⃣ Convert Struct to JSON (Serialization)
func structToJson() {
	person := Person{
		Name:   "John Doe",
		Age:    30,
		Email:  "johndoe@example.com",
		Active: true,
	}

	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}

	fmt.Println("Struct to JSON:")
	fmt.Println(string(jsonData))
}

// 3️⃣ Convert JSON to Struct (Deserialization)
func jsonToStruct() {
	jsonString := `{
		"name": "Alice",
		"age": 25,
		"email": "alice@example.com",
		"active": true
	}`

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}

	fmt.Println("\nJSON to Struct:")
	fmt.Printf("Name: %s, Age: %d, Email: %s, Active: %v\n",
		person.Name, person.Age, person.Email, person.Active)
}

// 4️⃣ Write JSON to File
func writeJsonToFile(fileName string) {
	person := Person{
		Name:   "Bob",
		Age:    28,
		Email:  "bob@example.com",
		Active: true,
	}

	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}

	err = ioutil.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
	fmt.Println("\nJSON data written to", fileName)
}

// 5️⃣ Read JSON from File
func readJsonFromFile(fileName string) {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var person Person
	err = json.Unmarshal(fileData, &person)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}

	fmt.Println("\nJSON read from file:")
	fmt.Printf("Name: %s, Age: %d, Email: %s, Active: %v\n",
		person.Name, person.Age, person.Email, person.Active)
}

// 6️⃣ Handle JSON Arrays
func handleJsonArray() {
	jsonArray := `[{"name":"Alice","age":24,"email":"alice@example.com","active":true},
				   {"name":"Bob","age":28,"email":"bob@example.com","active":false}]`

	var people []Person
	err := json.Unmarshal([]byte(jsonArray), &people)
	if err != nil {
		fmt.Println("Error deserializing JSON array:", err)
		return
	}

	fmt.Println("\nJSON Array to Struct Slice:")
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d, Email: %s, Active: %v\n",
			person.Name, person.Age, person.Email, person.Active)
	}
}

func main() {
	// 2️⃣ Struct to JSON (Serialization)
	structToJson()

	// 3️⃣ JSON to Struct (Deserialization)
	jsonToStruct()

	// 4️⃣ Write JSON to File
	writeJsonToFile("person.json")

	// 5️⃣ Read JSON from File
	readJsonFromFile("person.json")

	// 6️⃣ Handle JSON Arrays
	handleJsonArray()
}
