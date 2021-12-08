package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	ch := make(chan string)

	go func() {
		defer wg.Done()
		d := <-ch
		fmt.Println("Child: recevied signal:", d)
	}()

	ch <- "data"
	fmt.Println("Parent: sent signal")

	wg.Wait()
	fmt.Println("program complete")
}