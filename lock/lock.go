package main

import "sync"

var rwMutex sync.RWMutex
var sharedData map[string]int

func main() {
	sharedData = make(map[string]int)

	// 写入操作
	go func() {
		rwMutex.Lock() // 加写锁
		sharedData["key1"] = 1
		rwMutex.Unlock() // 解写锁
	}()

	// 读取操作
	//go func() {
	//rwMutex.RLock() // 加读锁
	//value := sharedData["key1"]
	//rwMutex.RUnlock() // 解读锁
	// ...
	//}()
}
