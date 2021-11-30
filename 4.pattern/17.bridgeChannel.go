package main

import "fmt"

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})

		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:  // Channel that you're reading from might have been canceled, so using read operator, prevent goroutine leaks.
					if ok == false {
						return
					}

					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()

		return valStream
	}
	// <-chan <-chan interface{} is a sequence of channels, can be read twice
	bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)

			for {
				var stream <-chan interface{}

				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}

				// When the stream we're currently looping over is closed,
				// we break out the loop performing the read from this 2.channel,
				// and continue with the next iteration of loop,
				// selecting channels to read from.
				// This provides us with an unbroken stream of values.
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()

		return valStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))

		go func() {
			defer close(chanStream)

			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()

		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}

// In some circumstances, you may find yourself wanting to consume values
// from a sequence of channels. One example might be a pipeline stage
// whose lifetime is intermittent.