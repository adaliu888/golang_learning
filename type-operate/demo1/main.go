package main

import "fmt"

// 定义一个接受变长参数的函数
func sum(args ...int) {
	total := 0
	for _, value := range args {
		total += value
	}
	fmt.Println("Sum:", total)
}

func main() {
	// 调用函数，传递不同数量的参数
	sum(1, 2, 3) // 输出: Sum: 6
	sum(10, 20)  // 输出: Sum: 30
	sum(42)      // 输出: Sum: 42
	sum()        // 输出: Sum: 0 (没有传递参数)
}
