package main

import "fmt"

func main() {
	fmt.Println("hello world")
	ch1 := make(chan int, 1)
	go func() {
		ch1 <- 1
	}()

	fmt.Println(<-ch1)
	close(ch1)     // close channel
	v, ok := <-ch1 // channel is closed,so this is default when call again
	if !ok {       //predicate channel whe
		fmt.Println("ch1 closed")
		panic(v) //panic:0 if channel is closed,panic is 0.
	}

}
