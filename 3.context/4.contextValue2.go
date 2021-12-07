package main

import (
	"context"
	"fmt"
)

func main()  {
	type user struct {
		name string
	}

	type userKey int

	u := user{name: "Andrew"}

	const uk userKey = 0

	ctx := context.WithValue(context.Background(), uk, &u)

	if u, ok := ctx.Value(uk).(*user); ok {
		fmt.Println("User", u.name)
	}

	if _, ok := ctx.Value(0).(*user); !ok {
		fmt.Println("User not found")
	}
}
