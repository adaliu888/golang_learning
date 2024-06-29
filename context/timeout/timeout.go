package main

import (
	context "context"
	"fmt"
	"sync"
	time "time"
)

func main() {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	// do something for 5 seconds
	go func() {
		fmt.Println("start do something")
		time.Sleep(2 * time.Second)
		fmt.Println("end do something")
	}()

	// check if context is done
	//// 使用WaitGroup等待goroutine完成或超时。
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("the operate is cancelled because of timeout")
		case <-time.After(10 * time.Second):
			fmt.Println("2 seconds later to wait for the task to finish")

		}
	}()
	// 等待goroutine完成或超时
	wg.Wait()
	fmt.Println("end")
}
