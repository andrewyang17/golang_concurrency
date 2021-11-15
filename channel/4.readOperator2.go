package main

import "fmt"

func main() {
	stringStream := make(chan int)
	close(stringStream)

	// Read from a closed channel
	salutation, ok := <- stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)
}

// We still able to perform a read operation,
// we could continue performing reads on this channel indefinitely
// despite the channel remaining closed.

// This is to allow support for multiple downstream reads from
// a single upstream writer on the channel.
// When you closed a channel, the downstream reads still available for it and
// other writers as to ensure the downstream keep on going, otherwise the program
// will just die, this is not the design we want.
// However, it opens up patterns for us.