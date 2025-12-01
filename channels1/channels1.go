package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	var sg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the lower limit: ")
	if !scanner.Scan() {
		log.Fatal("No input")
	}
	lowerLimit, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Invalid number")
		lowerLimit = 0
	}

	fmt.Print("Enter the upper limit: ")
	if !scanner.Scan() {
		log.Fatal("No input")
	}
	upperLimit, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Invalid number")
		upperLimit = 10000
	}

	startTime := time.Now()
	fmt.Println("started at: ", startTime)
	inputList := []int{}

	for i := lowerLimit; i <= upperLimit; i++ {
		inputList = append(inputList, i)
	}

	listLength := len(inputList)
	numWorkers := runtime.NumCPU()
	results := make(chan int, 150) //this will be read in main, channels are for communication and not for storage, hence large buffer size for result doesn't make sense.
	jobs := make(chan int, listLength)

	if listLength > numWorkers {
		numWorkers = listLength / 4 / numWorkers
	}
	if numWorkers > 5000 {
		numWorkers = 5000
	}
	for i := 0; i < numWorkers; i++ {
		sg.Add(1)
		go worker(jobs, results, &sg)
	}

	go func() {
		sg.Wait()
		close(results)
	}()

	for _, v := range inputList {
		jobs <- v
	}
	close(jobs)

	go func() {
		for range results {
			// discard
		}
	}()

	sg.Wait()
	endTime := time.Since(startTime)
	fmt.Println("ended at: ", endTime)
	fmt.Print("No of go routines used for the job:", numWorkers)
}

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range jobs {
		results <- item * item
	}
}
