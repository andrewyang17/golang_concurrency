package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	children := 2000
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(3)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)

				ch <- "data"
				fmt.Println("Child: sent signal")
			}
			<-sem
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println("Parent: received signal:", d)
	}

	fmt.Println("program complete")
}

// Fan out/in semaphore pattern provides a mechanic to control the number of goroutines
// executing work at any given time.