// Struct.go
// A complete guide to structs in Go.

package main

import (
	"fmt"
)

// 1️⃣ Defining a Struct
type Person struct {
	Name string
	Age  int
}

// 2️⃣ Struct with Methods
type Rectangle struct {
	Width  float64
	Height float64
}

// Method to calculate area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method to calculate perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3️⃣ Struct with Embedded Fields
type Address struct {
	City  string
	State string
	Zip   string
}

type Employee struct {
	Person  // Embedded struct
	Job     string
	Salary  float64
	Address // Embedded struct
}

// 4️⃣ Struct with Tags
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// 5️⃣ Struct Constructor Function
type Car struct {
	Make  string
	Model string
	Year  int
}

// Constructor for Car struct
func NewCar(make, model string, year int) Car {
	return Car{
		Make:  make,
		Model: model,
		Year:  year,
	}
}

func main() {
	// 1️⃣ Creating Struct Instances
	var p1 Person
	p1.Name = "John"
	p1.Age = 25
	fmt.Println("Person 1:", p1)

	// Shorter Declaration
	p2 := Person{Name: "Alice", Age: 30}
	fmt.Println("Person 2:", p2)

	// 2️⃣ Using Struct Methods
	rect := Rectangle{Width: 5.0, Height: 10.0}
	fmt.Println("Rectangle Area:", rect.Area())
	fmt.Println("Rectangle Perimeter:", rect.Perimeter())

	// 3️⃣ Struct with Embedded Fields
	emp := Employee{
		Person:  Person{Name: "Robert", Age: 35},
		Job:     "Software Engineer",
		Salary:  75000,
		Address: Address{City: "New York", State: "NY", Zip: "10001"},
	}
	fmt.Println("Employee Details:", emp)
	fmt.Println("Employee Name:", emp.Name) // Accessing embedded field directly
	fmt.Println("Employee City:", emp.Address.City)

	// 4️⃣ Struct with Tags (Demonstration)
	u := User{
		Username: "tech_guru",
		Email:    "guru@example.com",
		Age:      28,
	}
	fmt.Println("User with Tags:", u)

	// 5️⃣ Using Constructor Function
	car := NewCar("Toyota", "Corolla", 2022)
	fmt.Println("Car Details:", car)

	// 6️⃣ Anonymous Struct
	book := struct {
		Title  string
		Author string
		Pages  int
	}{
		Title:  "Go Programming",
		Author: "John Doe",
		Pages:  250,
	}
	fmt.Println("Book Details:", book)

	// 7️⃣ Pointer to a Struct
	ptr := &Person{Name: "Diana", Age: 29}
	fmt.Println("Pointer Person Name:", ptr.Name)
	ptr.Age = 30
	fmt.Println("Updated Age via Pointer:", ptr.Age)

	// 8️⃣ Comparing Structs
	p3 := Person{Name: "John", Age: 25}
	fmt.Println("p1 equals p3:", p1 == p3)

	// 9️⃣ Array of Structs
	products := [2]Car{
		{"Ford", "Focus", 2018},
		{"Honda", "Civic", 2020},
	}
	fmt.Println("Product List:", products)

	// 🔟 Slice of Structs
	players := []Person{
		{Name: "Tom", Age: 22},
		{Name: "Jerry", Age: 24},
	}
	fmt.Println("Players List:")
	for _, player := range players {
		fmt.Println(player.Name, "-", player.Age)
	}
}
