package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("moving on")

	case <-ctx.Done():
		fmt.Println("work complete")
	}
}