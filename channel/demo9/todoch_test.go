package todoch_test

import (
	ch "golang_learning/channel/demo9"
	"testing"
)

func BenchmarkToDoChannel(b *testing.B) {
	b.Run("single-threaded", func(b *testing.B) {
		ch.ToDoChannel()

	})
}
