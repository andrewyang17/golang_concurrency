package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://www.ardanlabs.com/blog/post/index.xml", nil)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	tr := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := http.Client{
		Transport: &tr,
	}

	ch := make(chan error, 1)
	go func() {
		log.Println("Starting request")
		resp, err := client.Do(req)
		if err != nil {
			ch <- err
			return
		}

		defer resp.Body.Close()

		io.Copy(os.Stdout, resp.Body)
		ch <- nil
	}()

	select {
	case <-ctx.Done():
		log.Println("timeout, cancel work...")
	case err := <-ch:
		if err != nil {
			log.Println(err)
		}
	}
}