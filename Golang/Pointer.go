// Pointer.go
// A complete guide to pointers in Go.

package main

import (
	"fmt"
)

// 1️⃣ Basic Pointer Example
func basicPointer() {
	a := 10
	ptr := &a // Pointer to variable 'a'
	fmt.Println("Value of a:", a)
	fmt.Println("Pointer to a:", ptr)
	fmt.Println("Value via pointer:", *ptr) // Dereferencing

	// Modifying value using pointer
	*ptr = 20
	fmt.Println("Updated value of a:", a)
}

// 2️⃣ Passing Pointers to Functions
func updateValue(val *int) {
	*val = *val + 10
}

func pointerFunctionExample() {
	num := 50
	fmt.Println("Before update:", num)
	updateValue(&num) // Passing pointer
	fmt.Println("After update:", num)
}

// 3️⃣ Pointers with Structs
type Person struct {
	Name string
	Age  int
}

func modifyPerson(p *Person) {
	p.Name = "Updated Name"
	p.Age += 5
}

func structPointerExample() {
	p := Person{Name: "John", Age: 30}
	fmt.Println("Before modification:", p)
	modifyPerson(&p)
	fmt.Println("After modification:", p)
}

// 4️⃣ Pointer to Pointer (Double Pointer)
func doublePointerExample() {
	a := 42
	p := &a
	pp := &p // Pointer to pointer

	fmt.Println("Value of a:", a)
	fmt.Println("Pointer p:", p)
	fmt.Println("Value via p:", *p)
	fmt.Println("Pointer pp:", pp)
	fmt.Println("Value via pp:", **pp) // Double dereferencing
}

// 5️⃣ Array and Pointers
func arrayPointerExample() {
	arr := [3]int{10, 20, 30}
	ptr := &arr[0]
	fmt.Println("First element via pointer:", *ptr)

	// Pointer arithmetic (not directly supported, but via slices)
	slice := arr[:]
	for i := range slice {
		fmt.Printf("Element %d via pointer: %d\n", i, slice[i])
	}
}

// 6️⃣ Slice and Pointers
func slicePointerExample() {
	slice := []int{1, 2, 3}
	ptr := &slice[1] // Pointer to second element
	fmt.Println("Second element via pointer:", *ptr)

	// Modifying via pointer
	*ptr = 20
	fmt.Println("Modified slice:", slice)
}

// 7️⃣ Pointer with Maps
func mapPointerExample() {
	m := map[string]*int{}
	val := 100
	m["key1"] = &val
	fmt.Println("Value at key1:", *m["key1"])

	// Changing value through pointer
	*m["key1"] = 200
	fmt.Println("Updated value at key1:", *m["key1"])
}

// 8️⃣ Pointer Receiver Methods
type Rectangle struct {
	Width  float64
	Height float64
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func methodPointerExample() {
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Before scaling:", rect)
	rect.Scale(2)
	fmt.Println("After scaling:", rect)
}

// 9️⃣ Nil Pointers
func nilPointerExample() {
	var ptr *int
	fmt.Println("Value of nil pointer:", ptr)

	if ptr == nil {
		fmt.Println("Pointer is nil, assigning value...")
		value := 500
		ptr = &value
	}

	fmt.Println("Value after assignment:", *ptr)
}

// 🔟 Dangling Pointers (Go automatically handles memory, no dangling pointers)

func main() {
	fmt.Println("\n🔹 Basic Pointer Example:")
	basicPointer()

	fmt.Println("\n🔹 Passing Pointers to Functions:")
	pointerFunctionExample()

	fmt.Println("\n🔹 Pointers with Structs:")
	structPointerExample()

	fmt.Println("\n🔹 Pointer to Pointer (Double Pointer):")
	doublePointerExample()

	fmt.Println("\n🔹 Array and Pointers:")
	arrayPointerExample()

	fmt.Println("\n🔹 Slice and Pointers:")
	slicePointerExample()

	fmt.Println("\n🔹 Pointer with Maps:")
	mapPointerExample()

	fmt.Println("\n🔹 Pointer Receiver Methods:")
	methodPointerExample()

	fmt.Println("\n🔹 Nil Pointers:")
	nilPointerExample()
}
