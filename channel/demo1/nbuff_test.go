package nbuff_test

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkNBuff(b *testing.B) {
	ch := make(chan bool)
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		go func() {
			buf.WriteString("hello, world")
			ch <- true
		}()
	}
	for range ch {

		buf.Reset()
	}
	fmt.Println("BenchmarkNBuff finished")

}
