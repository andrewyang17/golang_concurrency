package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings {  // read from a nil channel => block
				fmt.Println(s)
			}
		}()

		return completed
	}

	doWork(nil)
	fmt.Println("Done.")
}

// Result: Done.

// In this example, the lifetime of the process is very short,
// but in a ream program, goroutines could easily be started at the beginning
// of a long-lived program. In the worse case, the main goroutines could continue
// to spin up goroutines throughout its life, causing creep in memory utilization.

// The way to mitigate this is to establish a signal between the parent goroutine
// and its children that allows the parent to signal cancellation to its children.