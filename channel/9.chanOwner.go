package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {  // Consumer only has access to read channel
		resultStream := make(chan int, 5)  // Lexical Confinement, prevent other goroutines from writing it
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")
}

// Always keep the scope of channel ownership small
// so that these things remain obvious.

// If your program introduced deadlocks or panic,
// I think you'll find that the scope of your channel ownership
// has either gotten too large, or ownership has become unclear.