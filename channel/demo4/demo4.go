package main

import "time"

func main() {
	time.Sleep(3 * time.Second)
	println("end5")
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
	select { //usage
	case <-ch1:
		println("ch1")
	case <-ch2:
		println("ch2")
	}

}
