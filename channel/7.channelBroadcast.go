// Example of using channel to rebuild the program in sync/newCondBroadcast.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	begin := make(chan interface{})

	subscribe := func(fn func()) {
		go func() {
			<-begin
			fn()
		}()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	time.Sleep(2*time.Second)
	close(begin)

	clickRegistered.Wait()
}
