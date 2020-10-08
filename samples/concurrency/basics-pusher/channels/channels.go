package main

import "fmt"

func main() {
	// make a new channel
	messages := make(chan string)
	go func() {
		// write to channel
		messages <- "Hello World!"
	}()

	// read from channel
	message := <-messages

	fmt.Printf("I got the message %s\n", message)
}
