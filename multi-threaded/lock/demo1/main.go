package main

import (
	"fmt"
	"sync"
)

func safeIncrement(wg *sync.WaitGroup, mutex *sync.Mutex, counter *int) {
	defer wg.Done()      // 确保函数执行完毕,计数减一
	defer mutex.Unlock() // 确保在函数退出时释放锁

	mutex.Lock() // 获取锁
	*counter++
	fmt.Printf("Counter is now %d\n", *counter)
}

func main() {
	var wg sync.WaitGroup //define sync.WaitGroup
	var mutex sync.Mutex  //define sync.Mutex
	counter := 0

	for i := 0; i < 5; i++ { // 5个goroutine
		wg.Add(1)
		go safeIncrement(&wg, &mutex, &counter)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
