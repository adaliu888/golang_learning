package todocontext_test

import (
	ct "golang_learning/channel/demo11"
	"testing"
	"time"
)

func TestToDoContext(t *testing.T) {
	// 记录函数执行前的调用堆栈
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("function panicked: %v", r)
		}
	}()

	// 调用函数
	startTime := time.Now()
	ct.ToDoContext()
	duration := time.Since(startTime)

	// 验证函数执行时间是否在预期范围内
	if duration > 3*time.Second {
		t.Errorf("function took too long to execute: %v", duration)
	}
}
