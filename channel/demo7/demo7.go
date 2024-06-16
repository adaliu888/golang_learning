package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)            //type is int
	ch2 := make(chan float32)        //type is float32
	ch3 := make(chan string)         //type is string
	ch4 := make(chan bool)           //type is bool
	ch5 := make(chan []int)          //type is []int list slice
	ch6 := make(chan map[string]int) //type is map, map[string]int
	ch7 := make(chan complex64)      //type is complex64
	type myPerson struct {
		name string
		age  int
	}
	ch8 := make(chan myPerson)       //type is struct,and struct is nil
	ch9 := make(chan *int)           //type is pointer to int
	ch10 := make(chan chan int)      //type is channel int
	ch11 := make(chan chan chan int) //type is reduce channel int
	ch12 := make(chan func())        //type is function

	ch13 := make(chan interface{}) //type is interface

	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 1.1
	}()
	go func() {
		ch3 <- "hello"
	}()

	go func() {
		ch4 <- true
	}()
	go func() {
		ch5 <- []int{1, 2, 3}
	}()
	go func() {
		ch6 <- map[string]int{"a": 1, "b": 2}
	}()
	go func() {
		ch7 <- complex64(1)
	}()
	go func() {
		ch8 <- myPerson{name: "hello", age: 1}
	}()
	go func() {
		var x int
		ch9 <- &x

	}()
	go func() {
		ch10 <- ch1
	}()
	go func() {
		ch11 <- ch10
	}()
	go func() {
		ch12 <- func() {
			fmt.Println("hello")
		}
	}()
	go func() {
		ch13 <- "hello world"
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)
	fmt.Println(<-ch4)
	fmt.Println(<-ch5)
	fmt.Println(<-ch6)
	fmt.Println(<-ch7)
	fmt.Println(<-ch8)
	fmt.Println(<-ch9)
	fmt.Println(<-ch10)
	fmt.Println(<-ch11)
	f := <-ch12
	f()
	fmt.Println(<-ch13)

}
