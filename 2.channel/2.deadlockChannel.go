package main

import "fmt"

func main() {
	stringStream := make(chan string)

	// When this anonymous goroutine exits,
	// Go correctly detects that all goroutines are asleep,
	// and reports a deadlock.
	go func() {
		if 0 != 1 {  // To ensure stringStream 2.channel never gets a value placed upon it
			return
		}
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}
