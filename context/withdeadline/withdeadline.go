package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// create the context
	basectx := context.Background()
	// 设置截止时间为当前时间加上5秒
	d := time.Now().Add(10 * time.Second)
	// 创建一个新的context，继承 basectx，并设置 deadline 为 d
	ctx, cancel := context.WithDeadline(basectx, d)
	defer cancel() // 只有在函数返回前才会执行 cancel
	//运行一个goroutine,来模拟长时间运行
	// 任务在 5 秒内被执行，并在 10 秒后被取消
	// 如果任务在 5 秒内被执行，ctx.Done() channel 将被关闭，任务将在 5 秒后被取消。
	// 如果任务在 10 秒后被执行，ctx.Done() channel 被关闭，任务将在 10 秒后被取消。
	// 任务在 10 秒后被取消，ctx.Done() channel 被关闭，任务将在 10 秒后被取消。
	// 任务在 10 秒后被取消，ctx.Done() channel 被关闭，任务将在 10 秒后被取消。
	// 任务在 10 秒后被取消，ctx.Done() channel 被关闭，任务将在 10 秒后被取消
	go func() {
		select {
		case <-ctx.Done(): // 如果 context 被取消或者到达 deadline，ctx.Done() channel 将被关闭
			log.Println("任务在截止时间后被取消")
			return
		default:
			log.Println("任务在 5 秒内被执行")
			time.Sleep(5 * time.Second) // 模拟任务执行时间
		}
	}()

	// 等待 5 秒
	time.Sleep(10 * time.Second)

}
