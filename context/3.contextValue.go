package main

import (
	"context"
	"fmt"
)

type ctxKey int

// We don't export the keys that to store the data in context. We use functions to cast and retrieve them.
const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func main() {
	ProcessRequest("andrewyang17", "abc123")
}

func UserID(ctx context.Context) string {
	return ctx.Value(ctxUserID).(string)
}

func AuthToken(ctx context.Context) string {
	return ctx.Value(ctxAuthToken).(string)
}

func ProcessRequest(userID, authToken string) {
	// Context's key and value are defined as interface{}, we lose Go's type safety when attempting to retrieve values.
	// Therefore, Go authors recommend that to define a custom key-type in your package.
	// In this program, we defined ctxUserID and ctxAuthToken.
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("Handling response for %v (auth: %v)\n", UserID(ctx), AuthToken(ctx))
}

// The larger issue is definitely the nature of what developers should store in instances of Context.
// Context package: use context values only for "request-scoped data" that "transits processes and API boundaries",
//                  not for passing optional parameters to functions.

// Hueristics from the author (Concurrency in Go)
// 1. The data should transit process or API boundaries.
// 2. The data should be immutable.
// 3. The data should trend toward simple types.
// 4. The data should be data, not types with methods.
// 5. The data should help decorate operations, not drive them.

// Another dimension to consider is how many layers this data might need to traverse before utilization.
// What is acceptable on one team may not be acceptable on another.
