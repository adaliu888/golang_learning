package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
//sync.Mutex是读锁
func main() {
	lock.Lock()
	fmt.Println("hello")
	lock.Unlock()
}
