//subcontext

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//创建一个父context
	baseCtx := context.Background()
	//创建一个子context，继承父context，并设置一个10s的timeout
	ctx, cancel := context.WithTimeout(baseCtx, 10*time.Second)
	defer cancel() // 只有在函数返回前才会执行 cancel

	// 创建一个goroutine来模拟一个任务
	go func(ctx context.Context) {
		select {
		case <-ctx.Done(): // 若ctx被取消，ctx.Done()会返回一个channel，可以被select来接收
			fmt.Println("任务被取消")
		case <-time.After(5 * time.Second): // 5s后，向ctx.Done()发送一个值
			fmt.Println("任务执行完成")
		}
	}(ctx)

	// 主函数����15s，保证任务可以正常执行
	time.Sleep(15 * time.Second)
}

func FetchTime(ctx context.Context) (time.Time, error) {
	time.After(15 * time.Second)
	return time.Now(), nil
}
