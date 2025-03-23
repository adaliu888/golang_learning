package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func main() {
	lock.Lock()
	fmt.Println("hello")
	lock.Unlock()
}
