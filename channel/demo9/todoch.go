// 通过range 遍历 channel
package todoch

import "fmt"

func ToDoChannel() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
