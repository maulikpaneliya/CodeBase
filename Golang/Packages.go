	// Packages.go
// A complete guide to packages and imports in Go.

package main

import (
    "fmt"      // Standard library package for formatted I/O
    "math"     // Standard library package for mathematical operations
    "strings"  // Standard library package for string manipulation
    "time"     // Standard library package for date and time
)

// 1️⃣ Custom Package Example
// To use custom packages, create a folder structure like:
// myproject/
// ├── main.go
// └── greetings/
//     └── greetings.go

// greetings/greetings.go
// package greetings
//
// import "fmt"
//
// func SayHello(name string) {
//     fmt.Println("Hello,", name)
// }

// main.go (In the same myproject directory)
// package main
//
// import (
//     "fmt"
//     "myproject/greetings" // Importing custom package
// )
//
// func main() {
//     greetings.SayHello("Alice")
// }

func main() {
    // 2️⃣ Using Standard Library Packages

    // fmt package - Printing output
    fmt.Println("Welcome to Go Packages!")

    // math package - Mathematical operations
    fmt.Println("Square root of 16:", math.Sqrt(16))
    fmt.Println("Power of 2^3:", math.Pow(2, 3))

    // strings package - String manipulations
    sample := "Go Programming"
    fmt.Println("Uppercase:", strings.ToUpper(sample))
    fmt.Println("Contains 'Go':", strings.Contains(sample, "Go"))
    fmt.Println("Replace 'Go' with 'Golang':", strings.ReplaceAll(sample, "Go", "Golang"))

    // time package - Working with time
    currentTime := time.Now()
    fmt.Println("Current Time:", currentTime.Format("2006-01-02 15:04:05"))

    // 3️⃣ Multiple Imports
    // Multiple packages can be imported using parentheses (import block).

    // 4️⃣ Alias for Imports
    importExample()

    // 5️⃣ Blank Identifier Import
    blankImportExample()
}

// 4️⃣ Import Alias Example
import (
    m "math" // Giving alias 'm' for math package
    s "strings"
)

func importExample() {
    fmt.Println("Using alias 'm':", m.Sqrt(25))
    fmt.Println("Using alias 's':", s.ToLower("ALIAS EXAMPLE"))
}

// 5️⃣ Blank Identifier (_) Import Example
// Blank import is used when we want to run the package's init() function without using any exported functions directly.

import _ "net/http" // Only runs the init() function of net/http package

func blankImportExample() {
    fmt.Println("Blank import example - net/http package init() executed (no direct use)")
}
