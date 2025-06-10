package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 基本的 goroutine 示例
func basicGoroutine() {
	fmt.Println("=== 基本 Goroutine 示例 ===")
	go func() {
		fmt.Println("这是一个 goroutine")
	}()
	time.Sleep(time.Second) // 等待 goroutine 完成
}

// 2. 使用 channel 进行 goroutine 间通信
func channelExample() {
	fmt.Println("\n=== Channel 通信示例 ===")
	ch := make(chan string)
	
	go func() {
		ch <- "通过 channel 发送消息"
	}()
	
	msg := <-ch
	fmt.Println("接收到的消息:", msg)
}

// 3. 带缓冲的 channel
func bufferedChannelExample() {
	fmt.Println("\n=== 带缓冲的 Channel 示例 ===")
	ch := make(chan int, 2) // 缓冲区大小为 2
	
	ch <- 1
	ch <- 2
	// ch <- 3 // 这里会阻塞，因为缓冲区已满
	
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}

// 4. select 多路复用
func selectExample() {
	fmt.Println("\n=== Select 多路复用示例 ===")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(time.Second)
		ch1 <- "来自 ch1 的消息"
	}()
	
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "来自 ch2 的消息"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

// 5. 使用 WaitGroup 等待多个 goroutine 完成
func waitGroupExample() {
	fmt.Println("\n=== WaitGroup 示例 ===")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器
		go func(id int) {
			defer wg.Done() // 完成时减少计数器
			fmt.Printf("Worker %d 开始工作\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d 完成工作\n", id)
		}(i)
	}
	
	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("所有工作完成")
}

// 6. 互斥锁示例
func mutexExample() {
	fmt.Println("\n=== Mutex 互斥锁示例 ===")
	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()   // 加锁
			counter++      // 临界区
			mutex.Unlock() // 解锁
		}()
	}
	
	wg.Wait()
	fmt.Println("最终计数:", counter)
}

// 7. 并发安全的 map
func syncMapExample() {
	fmt.Println("\n=== Sync.Map 示例 ===")
	var syncMap sync.Map
	
	// 存储数据
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")
	
	// 读取数据
	if value, ok := syncMap.Load("key1"); ok {
		fmt.Println("key1:", value)
	}
	
	// 删除数据
	syncMap.Delete("key2")
	
	// 遍历数据
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
}

func main() {
	// 运行所有示例
	basicGoroutine()
	channelExample()
	bufferedChannelExample()
	selectExample()
	waitGroupExample()
	mutexExample()
	syncMapExample()
} 