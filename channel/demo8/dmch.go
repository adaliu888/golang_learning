package dmch

import (
	"fmt"
)

func DmChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 2 个 goroutine 交替执行 select 语句 channel send 和 channel recv
	go func() {
		select {
		case ch1 <- 1:
			fmt.Println("ch1")
		case ch2 <- 2:
			fmt.Println("ch2")
		}
		fmt.Println("channel send")
	}()

	select {
	case <-ch1:
		fmt.Println(ch1) //channel is cap
	case <-ch2:
		fmt.Println(ch2)
	}

	fmt.Println("channel recv")

}
