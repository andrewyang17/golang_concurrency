package main

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

	done := make(chan interface{})
	defer close(done)

	myChan := make(chan interface{})
	for val := range orDone(done, myChan) {
		// Do something with val
	}
}

// orDone is a separate single goroutine that abstract away the verbosity,
// and helps us to avoid repetitive code and nasty code smell.

// Always try for readability first, and avoid premature optimization.

// It mitigates the inefficient construction code like below:
	// 1.
	//for val := range myChan {...}

	//2.
	// loop:
	// for {
	//     select {
	//     case <-done:
	//         break loop
	//     case maybeVal, ok := <-myChan:
	//        if ok == false {
	//            return
	//     }
	//     // Do something with val
	// }
