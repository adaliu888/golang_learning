package main

import "time"

func main() {
	time.Sleep(3 * time.Second) // sleep for 3 seconds
	println("end5")
	ch1 := make(chan int) //define a channel ch1
	ch2 := make(chan int) //define a channel ch2
	go func() {           // goruntine
		time.Sleep(1 * time.Second)
		ch1 <- 1 //send 1 message
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 2 //send 2 messages
	}()
	select { //use select to receive events
	case <-ch1:
		println("ch1")
	case <-ch2:
		println("ch2")
	}

}
