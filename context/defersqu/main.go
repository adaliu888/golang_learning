package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	ch := make(chan string, 4)

	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			ch <- "hello" + string(i)
		}(i)
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}

}
