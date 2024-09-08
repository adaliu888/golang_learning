package heap_test

import (
	h1 "container/heap"
	"golang_learning/calculate/heap"
	"testing"
)

// TestIntHeap 测试 IntHeap 是否正确实现了 heap.Interface
func TestHead(t *testing.T) {
	// 创建一个 IntHeap 实例

	h := &heap.IntHeap{1, 2, 3}

	// 保证 IntHeap 实现了 heap.Interface
	t.Log(h)

	h1.Init(h)
	h1.Push(h, 4)
	if got, want := h.Len(), 4; got != want {
		t.Errorf("IntHeap.Len() = %v, want %v", got, want)
	}

	if got, want := h1.Pop(h), 1; got.(int) != want {
		t.Errorf("IntHeap.Pop() = %v, want %v", got, want)
	}
	if got, want := h.Len(), 3; got != want {
		t.Errorf("IntHeap.Len() = %v, want %v", got, want)
	}

}
