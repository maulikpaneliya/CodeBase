// Variables.go
// A complete guide to declaring and using variables in Go.

package main

import "fmt"

func main() {
    // 1️⃣ Declaring Variables (Using var keyword)
    var name string
    var age int
    var salary float64
    var isEmployed bool

    // Assigning values
    name = "John"
    age = 30
    salary = 50000.75
    isEmployed = true

    // Displaying values
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Salary:", salary)
    fmt.Println("Employed:", isEmployed)

    // 2️⃣ Short Variable Declaration (:=)
    country := "USA"
    score := 95.5
    isPass := true

    fmt.Println("Country:", country)
    fmt.Println("Score:", score)
    fmt.Println("Passed:", isPass)

    // 3️⃣ Multiple Variable Declaration
    var x, y, z int = 10, 20, 30
    a, b := "Hello", "World"
    fmt.Println("x:", x, "y:", y, "z:", z)
    fmt.Println(a, b)

    // 4️⃣ Constants (Cannot be changed)
    const pi = 3.14159
    fmt.Println("Value of pi:", pi)

    // 5️⃣ Zero Values (Default values when not assigned)
    var num int
    var str string
    var flag bool

    fmt.Println("Default int:", num)    // 0
    fmt.Println("Default string:", str) // ""
    fmt.Println("Default bool:", flag)  // false

    // 6️⃣ Variable Scope
    {
        innerVar := "I am inside a block"
        fmt.Println(innerVar)
    }
    // fmt.Println(innerVar) // Error: out of scope

    // 7️⃣ Type Inference
    var greeting = "Welcome" // Automatically inferred as string
    fmt.Println(greeting)

    // 8️⃣ Type Conversion
    var num1 int = 100
    var num2 float64 = float64(num1) // int to float64
    var num3 int = int(num2)         // float64 to int

    fmt.Println("int to float64:", num2)
    fmt.Println("float64 to int:", num3)

    // 9️⃣ Constants with iota (Enumeration)
    const (
        Sun = iota
        Mon
        Tue
        Wed
        Thu
        Fri
        Sat
    )
    fmt.Println("Sunday:", Sun, "Monday:", Mon, "Tuesday:", Tue)

    // 🔚 End of Variables.go
}
