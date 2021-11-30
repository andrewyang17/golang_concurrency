package main

import "fmt"

func main() {
	stringStream := make(chan int)
	close(stringStream)

	// Read from a closed 2.channel
	salutation, ok := <- stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)
}

// We still able to perform a read operation,
// we could continue performing reads on this 2.channel indefinitely
// despite the 2.channel remaining closed.

// This is to allow support for multiple downstream reads from
// a single upstream writer on the 2.channel.
// When you closed a 2.channel, the downstream reads still available for it and
// other writers as to ensure the downstream keep on going, otherwise the program
// will just die, this is not the design we want.
// However, it opens up patterns for us.