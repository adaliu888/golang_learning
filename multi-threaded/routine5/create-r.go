package main

import (
	"fmt"
	"sync"
)

// 同步启动1个routine
func doWork(i int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doWork(i, &wg, ch)
	}
	close(ch)

	wg.Wait()

	//close(ch) //我们创建了一个 channel 并发送了一些数据。发送完毕后，我们使用 close 函数关闭了 channel。然后，我们使用 for 循环和 range 语句接收并打印 channel 中的所有数据。当 channel 中没有更多数据时，循环自然结束。尝试在 channel 关闭后再次关闭它将会导致 panic，因此在实际应用中应该避免这种情况

}
