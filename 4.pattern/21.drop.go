package main

import "fmt"

func main() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("Child: received signal:", p)
		}
	}()

	const work = 2000

	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("Parent: sent signal")
		default:
			fmt.Println("Parent: dropped data")
		}
	}

	close(ch)
	fmt.Println("Parent: sent shutdown signal")

	fmt.Println("program complete")
}

// Drop pattern is an important pattern for services that may experience heavy loads
// at times and can drop requests when the service reaches a capacity of pending
// requests.