/*
Problem: Prime Number Finder

Generate numbers from 1 to 1000.

Launch multiple worker goroutines to check if a number is prime.

Each worker sends the prime numbers it finds to a results channel.

Collect all primes in the main goroutine and print them sorted.

Constraints:

Use channels to send numbers to workers and collect results.

Use a WaitGroup to wait for all workers to finish.

Make sure the main goroutine collects results without deadlocks.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	// Get max number from user
	fmt.Print("Enter the max number: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	maxNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Invalid number")
	}

	numWorkers := runtime.NumCPU() * 30 // number of workers
	jobs := make(chan int, 100)         // channel for numbers to check
	results := make(chan int, 100)      // channel for primes

	var wg sync.WaitGroup

	// Launch workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Collector goroutine: closes results when workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Send numbers to workers
	for i := 1; i <= maxNumber; i++ {
		jobs <- i
	}
	close(jobs) // no more jobs

	// Collect and print primes
	for prime := range results {
		fmt.Println(prime)
	}
}

func worker(inputChannel <-chan int, resultChannel chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range inputChannel {
		if resultItem := isPrime(item); resultItem {
			resultChannel <- item
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
