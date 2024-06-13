package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("CURRENT TIME: ", time.Now().Year(), time.Now().Month(), time.Now().Day())

	ch1 := make(chan int, 2) //define a channel and buffer is 2
	ch1 <- 1
	ch1 <- 2
	//ch1 <- 3           //缓冲去已经满了，发生阻塞
	fmt.Println(<-ch1) //channel send and receive messages is parallel
	fmt.Println(<-ch1)
	//fmt.Println(<-ch1) //当接收为空，goroutine等待，发生阻塞

}
