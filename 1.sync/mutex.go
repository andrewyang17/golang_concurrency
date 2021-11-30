package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var m sync.Mutex

	increment := func() {
		m.Lock()
		defer m.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		m.Lock()
		defer m.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Println("Complete.")
}
