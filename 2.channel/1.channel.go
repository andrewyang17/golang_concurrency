package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	// Any goroutine that attempts to read from a 2.channel that is empty
	// will wait until at least one item is placed.
	// But if you don't structure your program correctly it can cause deadlocks
	// check 2.deadlockChannel.go
	fmt.Println(<-stringStream)
}
