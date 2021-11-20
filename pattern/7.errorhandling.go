package main

import (
	"fmt"
	"net/http"
)

func main() {
	type Result struct {
		Error error
		Response *http.Response
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result

				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}

				select {
				case <-done:
					return
				case results <-result:
				}
			}
		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

// TIP:
// If your goroutine can produce errors,
// those errors should be tightly coupled with your result type,
// and passed along through the same lines of communication.

// We allow our main goroutine to make decision about what to do
// when errors occur. We've successfully separated the concerns of
// error handling from our producer goroutine. This is desirable because
// the goroutine that spawned the producer goroutine has more context about
// the running program, and can make more intelligent decisions about what to do
// with errors.


// We can simply remove the usage of done channel
// as the program is still executable without it, but
// it would remove the clarity of the code,
// which is not recommended.