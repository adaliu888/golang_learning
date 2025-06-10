package main

import (
	"context"
	"fmt"
	"time"
)

// 1. 基本的 Context 传递
func basicContextPass() {
	fmt.Println("\n=== 基本 Context 传递示例 ===")
	// 创建根 Context
	ctx := context.Background()
	
	// 传递 Context 到函数
	processWithContext(ctx)
}

func processWithContext(ctx context.Context) {
	// 在函数中继续传递 Context
	go subProcessWithContext(ctx)
	
	// 使用 Context 的值
	if ctx.Value("key") != nil {
		fmt.Println("主处理中获取到值:", ctx.Value("key"))
	}
}

func subProcessWithContext(ctx context.Context) {
	// 子函数中也能访问到 Context
	if ctx.Value("key") != nil {
		fmt.Println("子处理中获取到值:", ctx.Value("key"))
	}
}

// 2. 带值的 Context 传递链
func valueContextChain() {
	fmt.Println("\n=== Context 值传递链示例 ===")
	// 创建带值的 Context
	ctx := context.WithValue(context.Background(), "requestID", "req-001")
	ctx = context.WithValue(ctx, "userID", "user-001")
	
	// 模拟处理链
	handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
	// 打印当前层级的请求信息
	fmt.Printf("处理请求 [%s] 用户 [%s]\n",
		ctx.Value("requestID"),
		ctx.Value("userID"))
	
	// 添加新的值并传递
	newCtx := context.WithValue(ctx, "stage", "processing")
	processStage(newCtx)
}

func processStage(ctx context.Context) {
	// 可以访问到所有之前传递的值
	fmt.Printf("处理阶段 [%s] 请求 [%s] 用户 [%s]\n",
		ctx.Value("stage"),
		ctx.Value("requestID"),
		ctx.Value("userID"))
}

// 3. 带取消的 Context 传递
func cancelContextChain() {
	fmt.Println("\n=== Context 取消传递链示例 ===")
	ctx, cancel := context.WithCancel(context.Background())
	
	// 启动多个处理函数
	go processWithCancel(ctx, "处理1")
	go processWithCancel(ctx, "处理2")
	go processWithCancel(ctx, "处理3")
	
	// 2秒后取消所有处理
	time.Sleep(2 * time.Second)
	fmt.Println("发送取消信号...")
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func processWithCancel(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: 收到取消信号\n", name)
			return
		default:
			fmt.Printf("%s: 正在处理...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 4. 带超时的 Context 传递
func timeoutContextChain() {
	fmt.Println("\n=== Context 超时传递链示例 ===")
	// 创建一个 3 秒超时的 Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// 启动处理链
	go processWithTimeout(ctx)
	
	// 等待处理完成或超时
	<-ctx.Done()
	fmt.Println("主函数收到完成信号:", ctx.Err())
}

func processWithTimeout(ctx context.Context) {
	// 创建子 Context，继承父 Context 的超时设置
	subCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	// 启动子处理
	go subProcessWithTimeout(subCtx)
	
	// 处理自己的任务
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("处理被取消或超时:", ctx.Err())
			return
		default:
			fmt.Println("处理中...")
			time.Sleep(1 * time.Second)
		}
	}
}

func subProcessWithTimeout(ctx context.Context) {
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("子处理被取消或超时:", ctx.Err())
			return
		default:
			fmt.Println("子处理中...")
			time.Sleep(800 * time.Millisecond)
		}
	}
}

// 5. 多层 Context 传递
func multiLevelContext() {
	fmt.Println("\n=== 多层 Context 传递示例 ===")
	// 创建基础 Context
	ctx := context.Background()
	
	// 添加第一层值
	ctx = context.WithValue(ctx, "level", "1")
	ctx = context.WithValue(ctx, "data", "base")
	
	// 传递到第一层处理
	level1Process(ctx)
}

func level1Process(ctx context.Context) {
	// 添加第二层值
	ctx = context.WithValue(ctx, "level", "2")
	ctx = context.WithValue(ctx, "data", "level1")
	
	fmt.Printf("Level 1: level=%s, data=%s\n",
		ctx.Value("level"),
		ctx.Value("data"))
	
	// 传递到第二层处理
	level2Process(ctx)
}

func level2Process(ctx context.Context) {
	// 添加第三层值
	ctx = context.WithValue(ctx, "level", "3")
	ctx = context.WithValue(ctx, "data", "level2")
	
	fmt.Printf("Level 2: level=%s, data=%s\n",
		ctx.Value("level"),
		ctx.Value("data"))
	
	// 传递到第三层处理
	level3Process(ctx)
}

func level3Process(ctx context.Context) {
	fmt.Printf("Level 3: level=%s, data=%s\n",
		ctx.Value("level"),
		ctx.Value("data"))
}

func main() {
	// 运行所有示例
	basicContextPass()
	valueContextChain()
	cancelContextChain()
	timeoutContextChain()
	multiLevelContext()
} 