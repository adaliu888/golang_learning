package main

import (
	"errors"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doWork() error {
	defer func() {
		fmt.Println("deferred func")
		wg.Done()
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	return errors.New("some error with deadlock")
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	ch1 <- 1
	ch2 <- "hello"

	wg.Add(1)

	go func() {
		select {
		case a, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 closed")
			} else {
				fmt.Println(a)
			}
		case b, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 closed")
			} else {
				fmt.Println(b)
			}
		}
		wg.Wait()
	}()

}
