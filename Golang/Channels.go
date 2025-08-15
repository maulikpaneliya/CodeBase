// Channels.go
// This file demonstrates Go's channel patterns and usage

package main

import (
	"fmt"
	"time"
)

// Basic channel operations
func basicChannelOperations() {
	// Create an unbuffered channel
	ch := make(chan string)

	// Send data in a goroutine
	go func() {
		ch <- "Hello from channel!"
	}()

	// Receive data
	msg := <-ch
	fmt.Println(msg)
}

// Buffered channel example
func bufferedChannel() {
	// Create a buffered channel with capacity 2
	ch := make(chan string, 2)

	// Send without blocking
	ch <- "First message"
	ch <- "Second message"

	// Receive messages
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Channel direction (send-only or receive-only)
func channelDirection() {
	ch := make(chan int)

	// Sender
	go sender(ch)

	// Receiver
	receiver(ch)
}

// Sender function - can only send to channel
func sender(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}

// Receiver function - can only receive from channel
func receiver(ch <-chan int) {
	// Range over channel until it's closed
	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}
}

// Select statement with channels
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// Wait for messages from either channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

// Timeout pattern using select
func timeoutPattern() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Response"
	}()

	select {
	case response := <-ch:
		fmt.Println("Received:", response)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}
}

func main() {
	fmt.Println("=== Basic Channel Operations ===")
	basicChannelOperations()

	fmt.Println("\n=== Buffered Channel ===")
	bufferedChannel()

	fmt.Println("\n=== Channel Direction ===")
	channelDirection()

	fmt.Println("\n=== Select Statement ===")
	selectExample()

	fmt.Println("\n=== Timeout Pattern ===")
	timeoutPattern()
}
