package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("Child: sent signal")
	}()

	d := <-ch
	fmt.Println("Parent: received signal:", d)

	fmt.Println("Program complete")
}