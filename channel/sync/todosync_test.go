package todosync_test

import (
	"fmt"
	sc "golang_learning/channel/sync"
	"sync"
	"testing"
	"time"
)

func BenchmarkTodoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sc.ToDoSync()
	}
}

func TestTodoSync(t *testing.T) {
	startTime := time.Now()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("function panicked: %v", r)
		}
	}()

	var wg sync.WaitGroup
	ch := make(chan int, 100)

	wg.Add(200) // 增加计数，因为有两个goroutine集合

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
	}

	wg.Wait() // 等待所有goroutine完成
	duration := time.Since(startTime)

	if duration > 3*time.Second {
		t.Errorf("function took too long to execute: %v", duration)
	}
}
