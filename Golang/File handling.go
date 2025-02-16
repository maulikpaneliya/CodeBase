// FileHandling.go
// A complete guide to file handling operations in Go.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 1️⃣ Create a File and Write Data
func createFile() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, this is a sample text.\nWelcome to Go file handling.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File created and data written successfully.")
}

// 2️⃣ Read File (Method 1: ioutil)
func readFileUsingIoutil() {
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("\nRead using ioutil:")
	fmt.Println(string(data))
}

// 3️⃣ Read File (Method 2: os and bufio)
func readFileUsingBufio() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("\nRead using bufio:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// 4️⃣ Append Data to File
func appendToFile() {
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("This line is appended to the file.\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("Data appended to file successfully.")
}

// 5️⃣ Copy File
func copyFile(src, dst string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	fmt.Println("File copied successfully from", src, "to", dst)
}

// 6️⃣ Rename File
func renameFile(oldName, newName string) {
	err := os.Rename(oldName, newName)
	if err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}
	fmt.Println("File renamed from", oldName, "to", newName)
}

// 7️⃣ Delete File
func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File", fileName, "deleted successfully.")
}

// 8️⃣ Check if File Exists
func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func main() {
	// 1️⃣ Create and Write File
	createFile()

	// 2️⃣ Read File using ioutil
	readFileUsingIoutil()

	// 3️⃣ Read File using bufio
	readFileUsingBufio()

	// 4️⃣ Append to File
	appendToFile()

	// 5️⃣ Copy File
	copyFile("example.txt", "copy_example.txt")

	// 6️⃣ Rename File
	if fileExists("copy_example.txt") {
		renameFile("copy_example.txt", "renamed_example.txt")
	}

	// 7️⃣ Delete File
	if fileExists("renamed_example.txt") {
		deleteFile("renamed_example.txt")
	}

	// Display final content of example.txt
	readFileUsingIoutil()
}
