package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)

		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)  // What else can it do? It can't pass back! Don't put your goroutines in this awkward position!
					continue
				}

				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()

		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://google.com", "http://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

// Always separate your concerns, in general, your concurrent processes should
// send their errors to another part of your program that has complete information
// about the state of your program, and make a more informed decision about what to do.