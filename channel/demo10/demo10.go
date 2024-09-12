package demo10

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	// 通过range 遍历 channel
	for v := range ch {
		switch v {
		case 0:

			fmt.Println("Received all values from channel")
			//break // 使用 break 退出循环

		default:
			fmt.Println("Value:", v)
		}
	}
}
