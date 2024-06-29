package main

import (
	"context"
	"fmt"
	"time"
)

//使用context.WithCancel，withtimeout 去取消channel，不会造成死锁；或者超时去取消channel，不会造成死锁

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if cancel != nil {
		fmt.Println("context canceled")
	}
	//do something
	ch := make(chan interface{})
	go func() {
		fmt.Println("start do something")
		time.Sleep(5 * time.Second)
		//ch <- 1 // send a message to cancel context
		fmt.Println("end do something")
	}()
	close(ch)
	for {
		select {
		case <-ch:
			fmt.Println("context canceled by channel")
			return
		default:
			fmt.Println("default case")
			time.Sleep(1 * time.Second)
		}
	}

}
