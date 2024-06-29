package main

import (
	"context"
	"time"
)

func main() {
	//create context,
	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个goroutine来模拟长时间运行的任务。
	go func() {

		// 启动一个goroutine来模拟长时间运行的任务。
		for {
			select {
			case <-ctx.Done(): // 如果context被取消，就退出循环
				return
			default:
				println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 等待5秒，让goroutine 运行一会
	time.Sleep(5 * time.Second)

	// 取消操作，这将触发goroutine中的取消逻辑。
	cancel() // 取消context ，触发通知子goroutine退出

	// 等待goroutine退出
	time.Sleep(5 * time.Second)

	go func() {

		// 启动一个goroutine来模拟长时间运行的任务。
		for {
			select {
			case <-ctx.Done(): // 如果context被取消，就退出循环
				return
			default:
				println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 等待5秒，让goroutine 运行一会
	time.Sleep(10 * time.Second)

	// 取消操作，这将触发goroutine中的取消逻辑。
	cancel() // 取消context ，触发通知子goroutine退出

	// 等待goroutine退出
	//time.Sleep(5 * time.Second)

	println("over")

}
