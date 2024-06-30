// mix context and waitgroup
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//create context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create waitgroup and channel to handle goroutines
	// channel has a buffer of 10 to prevent goroutine blocking
	// when all goroutines are finished and context is cancelled
	// we will receive an empty value from the channel in that case
	// and stop waiting for goroutines to finish
	// this is a simple example of context and waitgroup usage together,
	// in a real-world scenario you might want to use a more complex context or a context cancelation mechanism
	// such as context.WithTimeout or context.WithCancel to manage the context lifecycle more effectively
	// and avoid potential deadlocks or resource leaks
	// the context cancelation logic would be handled in a separate goroutine or in the main function
	// and would send a cancel signal to the context when the main function is done with its work
	// or when a specific timeout is reached.
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ch <- i

		go func(i int) {
			defer wg.Done()
			select {
			case data := <-ch:
				fmt.Printf("received %d from channel\n", data)
			case <-ctx.Done():
				fmt.Printf("goroutine %d was cancelled\n", i)
				return
			}
		}(i)
	} // for xun
	wg.Wait()
	fmt.Println("all goroutines finished")
	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("context cancelled")
	defer close(ch)
}
