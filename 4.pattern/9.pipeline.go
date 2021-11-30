package main

import (
	"fmt"
)

func main() {
	// Converts a discrete set of values into a stream of data on a 2.channel
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)

		go func() {
			defer close(intStream)

			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <-i:
				}
			}
		}()

	return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)

		go func() {
			defer close(multipliedStream)

			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i*multiplier:
				}
			}
		}()

		return multipliedStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)

		go func() {
			defer close(addedStream)

			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i+additive:
				}
			}
		}()

		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}

// Pipeline is nothing more than a series of stages that take data in,
// perform an operation on it, and pass the data back out.

// A stage consumes and returns the same type.

// This program turns out to have massive ramification,
// we'll discover in later Fan-Out, Fan-In.

// What would happen if we called close on the done 2.channel before
// the program was finished executing?
// It will force the pipeline stage to terminate.