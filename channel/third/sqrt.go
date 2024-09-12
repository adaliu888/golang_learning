package sqrt

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return float64(z)
	// 注意：在 Go 语言中，sqrt 并没有内置的函数，需要用循环来实现
}
