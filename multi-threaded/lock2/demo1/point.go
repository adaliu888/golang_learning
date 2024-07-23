package main

import (
	"fmt"
)

func main() {
	x := make(chan int)
	fmt.Println(x)
	close(x)
	fmt.Println(x)
	var s int = 10 //概念的模式

	var p *int
	fmt.Println(p)

	p = &s
	fmt.Println(p, *p)
	type Myinterface interface {
	}
	var mi2 Myinterface = &s
	var mi Myinterface
	fmt.Println(mi)
	fmt.Println(mi2) //从语义上有问题
	mi = mi2
	fmt.Println(mi)

}
