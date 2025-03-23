package main

//我们可以通过Tag来增强结构体的定义，Tag会带上一些meta信息，在本文中我们将通过几个例子来深入了解Tag的用法
//结构体,

import "fmt"

type T1 struct {
	f1 string
}
type T2 struct {
	T1
	f2     int64
	f3, f4 float64
}

func main() {
	t := T2{T1{"foo"}, 1, 2, 3}
	fmt.Println(t.f1)    // foo
	fmt.Println(t.T1.f1) // foo
	fmt.Println(t.f2)    // 1
	fmt.Println(t.f3)    // 2
	fmt.Println(t.f4)    // 3
}
