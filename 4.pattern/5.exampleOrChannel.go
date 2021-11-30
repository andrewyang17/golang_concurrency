package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {  // Base case
		case 0:  // Termination criteria
			return nil
		case 1:  // Second termination criteria
			return channels[0]
		}

		orDone := make(chan interface{})

		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:  // Every recursive call to or will at least have two channels
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()

		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-or(
		sig(3*time.Second),
		sig(6*time.Second),
		sig(9*time.Second),
		sig(12*time.Second),
		sig(15*time.Second),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}

// Our 2.channel that closes after 3 second causes the entire 2.channel created by the call
// to close. This is because, despite its place in the tree the or function builds -
// it will always close first and thus the channels that depend on its closure
// will close as well.

// There will be another way of doing this in the 3.context package.