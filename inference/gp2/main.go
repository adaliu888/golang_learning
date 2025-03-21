package main

import (
	"fmt"
	"sync"
)

// Worker function that processes jobs and sends results
func worker(id int, jobChan <-chan int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobChan {
		// Simulate some work
		result := job * job // Example: square the job number
		resultChan <- result
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
}

// CreatePool initializes a pool of workers
func createPool(poolSize int, jobChan <-chan int, resultChan chan<- int) {
	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go worker(i, jobChan, resultChan, &wg)
	}
	wg.Wait()
}

func main() {
	// Channels for jobs and results
	jobChan := make(chan int, 10)
	resultChan := make(chan int, 10)

	// Start the pool
	go createPool(64, jobChan, resultChan)

	// Send some jobs
	for i := 0; i < 100; i++ {
		jobChan <- i
	}
	close(jobChan)

	// Collect results
	for i := 0; i < 100; i++ {
		result := <-resultChan
		fmt.Printf("Result: %d\n", result)
	}
	close(resultChan)
}
