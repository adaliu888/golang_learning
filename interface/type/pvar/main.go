package main

import "fmt"

// 修改double函数以返回修改后的值
func double(x *int64) *int64 {
	*x *= 2  // 将原来的*x += *x改为*x *= 2，使数值翻倍
	return x // 返回指针
}

func main() {
	var a int64 = 3 // 将a的类型改为int64以匹配double函数的参数类型
	p := double(&a) // 将double函数的返回值赋给p

	fmt.Println(*p)          // 打印p指向的值
	fmt.Println(a)           // 打印a的值，它应该现在是6
	fmt.Println(a, p == nil) // 检查p是否为nil，它不应该为nil
}
