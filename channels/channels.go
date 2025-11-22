package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// So here, main go routine is writing to a channel and a child go routine is reading from it. Channels are blocking, it blocks the execution until someone on the other end is reading from it.
// You can push to buffered channels until their limit even if none is reading.

// Goal: Fetch the contents of multiple URLs concurrently and count the total number of bytes fetched.
func main() {
	startTime := time.Now()
	var wg sync.WaitGroup

	urls := []string{
		"https://seez.co",
		"https://starmark.dk",
		"https://nortonway.com",
		"https://www.google.com",
		"https://www.github.com",
	}

	resultChan := make(chan int, len(urls))
	for _, item := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			result := fetchContent(u)
			resultChan <- result
		}(item)
	}
	wg.Wait()
	close(resultChan)

	sum := 0
	for r := range resultChan {
		sum += r
	}
	fmt.Println(sum)
	endTime := time.Since(startTime)
	fmt.Println("Total time taken: " + endTime.String())
}

func fetchContent(url string) int {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return len(body)
	}
	return 0
}
