package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		// Closing a channel is one of the ways to signal MULTIPLE goroutines SIMULTANEOUSLY.
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	// Notice how the loop doesn't need an exit criteria.
	// The range doesn't return the second boolean value.
	// It exits only when the channel that being read is closed.
	for integer := range intStream {
		fmt.Println(integer)
	}

}

// If you have n goroutines waiting on a single channel,
// instead of writing n times to the channel to unblock each goroutine,
// you can simply close the channel.

// Since a closed channel can be read from an infinite number of times
// it doesn't matter how many goroutines are waiting on it,
// and closing the channel is both cheaper and faster than performing n writes.