package main

import (
	"fmt"
	"time"
)

// create a channel and not create buffer
func main() {
	ch := make(chan int) //no buffer
	go func() {
		time.Sleep(2 * time.Second) //sleep interval between 2 seconds
		ch <- 1                     //send a message
	}()

	val := <-ch //receive the message
	fmt.Println(val)
}
