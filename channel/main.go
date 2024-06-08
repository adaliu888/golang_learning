package main

import "fmt"

//create new channel
func main() {
	ch := make(chan int) //make a channel,deflaut not input the number of channel

	go func(a, b int) {
		sum := a + b

		ch <- sum //ch send message from var sum

	}(1, 2)

	r := <-ch //r get message from ch recive message

	fmt.Printf("computered value %v\n", r) //fmt.Printf 是 Go 语言标准库 fmt 包中的一个函数，它用于格式化并输出数据到 io.Writer 接口，通常是标准输出（控制台）。Printf 的行为类似于 C 语言中的 printf 函数，但它提供了更多的灵活性和安全性

}
