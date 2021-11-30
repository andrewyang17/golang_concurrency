package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			<-begin  // Wait until it's told it can continue, how? by closing the 2.channel
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)  // Unblocking all the goroutines simultaneously
	wg.Wait()
}
