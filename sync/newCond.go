package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remove from queue")
		c.L.Unlock()
		c.Signal()  // wakes one goroutine waiting on c
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			fmt.Println("length is equals to 2")
			// Wait doesn't just block, it suspends the current goroutine,
			// allowing other goroutines to run on the OS thread.
			// Upon entering Wait, Unlock is called on the Cond variable's Locker
			// and upon exiting Wait, Lock is called on the Cond variable's Locker
			c.Wait()
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second)
		c.L.Unlock()
	}
}
