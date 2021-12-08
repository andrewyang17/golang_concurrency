package main

import (
	"context"
	"fmt"
	"time"
)

func retryTimeout(ctx context.Context, retryInterval time.Duration,
	check func(ctx context.Context) error) {
	for {
		fmt.Println("Perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}

		fmt.Println("Check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired:", ctx.Err())
			return
		}

		fmt.Printf("Wait %s before trying again \n", retryInterval)
		t := time.NewTimer(retryInterval)

		select {
		case <-ctx.Done():
			fmt.Println("time expired:", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}

// Retry timeout pattern is great when want to ping something which might fail,
// but we don't want fail immediately, instead we want to retry for a specified
// amount of time before fail.