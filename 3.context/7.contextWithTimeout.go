package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	type data struct {
		UserID string
	}

	duration := 150 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan data, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <-data{"andrewyang17"}
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}

// Difference between with WithDeadline and WithTimeout is
// WithDeadline takes a time.Time type,
// WithTimeout takes a time.Duration,
// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).