package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doWork(i int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter when the goroutine completes
	fmt.Println("wait for all the goroutines to complete")
}

func main() {
	const N = 5 //number of goroutines
	for i := 0; i < N; i++ {
		wg.Add(1) // increment the counter when the goroutine starts
		go doWork(i, &wg)
	}
	wg.Wait() // wait for all the goroutines to complete
}
