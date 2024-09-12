/*
go语言中有一个其他的工具sync.WaitGroup 能更加方便的帮助我们达到这个目的。

WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。

使用WaitGroup 将上述代码可以修改为：
*/

package todosync

import (
	"fmt"
	"sync"
)

func ToDoSync() {
	var wg sync.WaitGroup
	ch := make(chan int, 100)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			ch <- 1
			fmt.Println("1")
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			<-ch
			fmt.Println("2")
			wg.Done()
		}()
		wg.Wait()
	}
}
