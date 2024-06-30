package main

// context 混合使用  waitgroup
import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan string, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	for {
		select {
		case <-ctx.Done():
			{
				fmt.Printf("Worker %d exiting\n", id)
				return
			}
		case get, ok := <-ch:
			{
				if ok {
					fmt.Printf("Worker %d received message: %s.\n", id, get)
				} else {
					fmt.Printf("Worker %d received channel closed.\n", id)
				}
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string, 2)

	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1)
	go worker(ctx, &wg, ch, 2)
	runtime.Gosched()
	wg.Add(1)
	go worker(ctx, &wg, ch, 1) //you need use go routine worker when  it is in the main func

	time.Sleep(1000 * time.Millisecond)
	runtime.Gosched()
	go func() {
		stra := [5]string{"a", "b", "c", "d", "e"}
		for _, ss := range stra {
			ch <- ss
			runtime.Gosched()
		}
	}()

	time.Sleep(time.Second * 2)
	cancel()
	wg.Wait()
}
