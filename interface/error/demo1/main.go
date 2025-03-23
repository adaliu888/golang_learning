package main

//返回值错误处理
import (
	"fmt"
	"math"
)

func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, fmt.Errorf("sqrt: negative number %g", i)
	}
	return math.Sqrt(i), nil
}

// 若sqrt函数返回一个错误，main函数会立即返回，不执行以下代码

func main() {
	var num float64 = 10
	_, err := sqrt(num)
	if err != nil {
		fmt.Println(err)
	}
}
