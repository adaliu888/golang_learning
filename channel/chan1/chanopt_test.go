package chanopt_test

import (
	CH "golang_learning/channel/chan1"
	"testing"
)

// BenchmarkChanOpt 是基准测试函数
func BenchmarkChanOpt(b *testing.B) {
	// 预热，确保在实际测量之前，所有的初始化和编译优化都已经完成
	b.Run("warmup", func(b *testing.B) {
		for i := 0; i < 5; i++ {
			CH.ChanOpt()
		}
	})

	b.ResetTimer() // 重置计时器，忽略预热的时间
	b.Run("progress", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CH.ChanOpt()
		}
	})

}
