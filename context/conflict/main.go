//not use select to conform channel is not supported

package main

import (
	"context"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 5)
		ch <- "done"
	}()
	close(ch)

	<-ctx.Done()
	println("在5秒内未完成，任务被取消。")

}
