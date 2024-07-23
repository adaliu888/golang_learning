package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doWork(i int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done() // decrement the counter when the goroutine completes
	select {
	case <-ctx.Done(): // check if the context has been cancelled
		fmt.Println("goroutine", i, "was cancelled")
		return // if cancelled, return from the goroutine immediately
	default:
		fmt.Println("wait for all the goroutines to complete", i)
	}

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	const N = 5 //number of goroutines
	for i := 0; i < N; i++ {
		wg.Add(1) // increment the counter when the goroutine starts
		go doWork(i, &wg, ctx)
	} // cancel the goroutines
	wg.Wait() // wait for all the goroutines to complete
	fmt.Println("all the goroutines have completed")
	cancel() // cancel the context, so the goroutines will be immediately cancelled
}
