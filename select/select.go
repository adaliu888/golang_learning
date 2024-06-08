package main

import (
	"time"
)

// create a main() function
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1

	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2
	}()
	// 使用select等待数据,if the data is already, then select the data to execute.
	select {
	case <-ch1:
		println("ch1")
	case <-ch2:
		println("ch2")
	}
	println("end")
	// 程序会在这里阻塞，直到ch1或ch2有数据
	time.Sleep(3 * time.Second)
	println("end2")
	// 程序会在这里阻塞，直到ch1或ch2有数据
	time.Sleep(3 * time.Second)
	println("end3")
	// 程序会在这里阻塞，直到ch1或ch2有数据
	time.Sleep(3 * time.Second)
	println("end4")
	// 程序会在这里阻塞，直到ch1或ch2有数据
	time.Sleep(3 * time.Second)
	println("end5")
	//close goroutine
	close(ch1)
	close(ch2)

}
