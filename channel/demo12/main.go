//withtimeout

// 1. 功能：WithTimeout函数可以让一个goroutine在特定的时间内执行任务，
// 如果在这个时间内任务没有完成，WithTimeout函数会返回一个error。
// 2. 场景：在需要对一个任务执行有一定时限的场景下，使用WithTimeout函数可以保证任务一定能完成。
// 3. 示例：
//      ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)	\
//      defer cancel()    \
//      // 执行任务并处理返回值
//      // ...
//      // 任务执行结束后，ctx.Done() channel 被关闭，从而可以检查任务是否被取消
//      if err := ctx.Err(); err != nil {
//          log.Println("任务在 5 秒内被执行")
//      } else {
//          log.Println("任务在 5 秒内被执行")
//      }

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 4)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("任务在 5 秒内被执行")
	case <-time.After(time.Second * 10):
		fmt.Println("任务在 10 秒内被执行")

	}
}
