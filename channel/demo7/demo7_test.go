package demo7_test

import (
	CH "golang_learning/channel/demo7"
	"testing"
)

func BenchmarkTDChannel(b *testing.B) {
	b.Run(
		"Channel",
		func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CH.TDChannel()
			}
		},
	)

}
