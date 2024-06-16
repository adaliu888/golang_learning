package main

import (
	"fmt"
	"time"
)

// select 进行多路复用，select 语句可以监听多个 channel，当其中一个 channel 中有数据时，select 语句将执行相应的 case 子句。select 语句用于多个channel操作之间继续宁选择
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	select { //select在多个channel中进行监听，实现多路复用
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	}

}
