package main

import (
	"fmt"
	"sync"
)

var rwMutex sync.RWMutex
var sharedData map[string]int

func main() {
	sharedData = make(map[string]int)
	sharedData := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
	}

	// 写入操作
	go func() {
		rwMutex.Lock() // 加写锁
		sharedData["key1"] = 1
		sharedData["key2"] = 2
		rwMutex.Unlock() // 解写锁
	}()

	//读取操作
	go func() {
		rwMutex.RLock() // 加读锁
		value := sharedData["key1"]
		fmt.Printf("value = %d\n", value)
		rwMutex.RUnlock() // 解读锁

	}()
}
