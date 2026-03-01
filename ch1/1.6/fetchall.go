package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // Create a "walkie-talkie" channel for strings

	for _, url := range os.Args[1:] {
		// The 'go' keyword spawns a new worker (Goroutine)
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		// Receive the messages from the walkie-talkie and print them
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Send error to channel
		return
	}

	// We don't want to print the HTML this time, we just count the bytes
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	// Send the final report back to the main program via the channel
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
