package main

import "fmt"

//main 函数创建了一个 bill 实例，设置了小费，并打印了格式化的账单
func main() {
	mybill := NewBill("mario is bill")

	fmt.Println(mybill.format())

}
