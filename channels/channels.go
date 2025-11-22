package main

import (
	"fmt"
	"strconv"
	"sync"
)

// So here, main go routine is writing to a channel and a child go routine is reading from it. Channels are blocking, it blocks the execution until someone on the other end is reading from it.
// You can push to buffered channels until their limit even if none is reading.
func main() {
	myChannel := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	go ReadChannel(myChannel, &wg)
	for i := 0; i < 3; i++ {
		myChannel <- "Hello " + strconv.Itoa(i)
	}
	close(myChannel)
	wg.Wait()
}

func ReadChannel(myChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range myChan {
		fmt.Println(item)
	}
}
