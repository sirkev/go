package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to communicate between goroutines
	ch := make(chan string)

	// Start two goroutines concurrently
	go printMessage("Hello", ch)
	go printMessage("World", ch)

	// Receive messages from the channel
	msg1 := <-ch
	msg2 := <-ch

	// Print the final message
	fmt.Println(msg1, msg2)
}

func printMessage(message string, ch chan<- string) {
	// Simulate some processing time
	time.Sleep(1 * time.Second)

	// Send the message to the channel
	ch <- message
}
