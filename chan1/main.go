package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	datachan := make(chan int) //make a channel ,but buffer does not infered

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				datachan <- doWork()
			}()
		}
		wg.Wait()
		close(datachan)
	}()

	for n := range datachan {
		fmt.Printf("n = %d\n", n)

	}
}
