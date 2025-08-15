// Goroutines.go
// This file demonstrates Go's concurrency features using goroutines

package main

import (
	"fmt"
	"sync"
	"time"
)

// Basic goroutine example
func basicGoroutine() {
	go func() {
		fmt.Println("Running in a goroutine")
	}()
	time.Sleep(time.Millisecond) // Wait for goroutine to complete
}

// WaitGroup example for synchronization
func waitGroupExample() {
	var wg sync.WaitGroup

	// Launch multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to complete
}

// Worker Pool implementation
func workerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

// Rate limiting example
func rateLimiting() {
	requests := make(chan int, 5)
	limiter := time.Tick(200 * time.Millisecond)

	// Send requests
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Process requests with rate limiting
	for req := range requests {
		<-limiter // Rate limit
		fmt.Printf("Processing request %d\n", req)
	}
}

func main() {
	fmt.Println("=== Basic Goroutine ===")
	basicGoroutine()

	fmt.Println("\n=== WaitGroup Example ===")
	waitGroupExample()

	fmt.Println("\n=== Worker Pool Example ===")
	workerPool()

	fmt.Println("\n=== Rate Limiting Example ===")
	rateLimiting()
}
