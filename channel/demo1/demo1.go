package main

import (
	"fmt"
)

//create a new channel instance with no buffering

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

}
