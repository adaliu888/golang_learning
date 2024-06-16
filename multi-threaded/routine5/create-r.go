package main

import (
	"fmt"
	"sync"
)

func doWork(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doWork(i, &wg)
	}
	wg.Wait()
}
