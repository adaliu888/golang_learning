package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	for v := range ch {
		fmt.Println(v)
	}

}
