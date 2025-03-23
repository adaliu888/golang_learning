package main

import (
	"fmt"
)

//两个结构体，使用定义一个基本类型，使用tag定义json的key不同，然后进行转换

type T1 struct {
	f int `json:"foo"`
}
type T2 struct {
	f int `json:"bar"`
}

func main() {
	t1 := T1{10}
	t2 := T2{20}
	fmt.Println(t1)
	fmt.Println(t2) // t2 as T2 type
	t2 = T2(t1)     //t1 as value for t2 of struct T2
	fmt.Println(t2) // {10}
}

//golang T1和T2是两个结构体，T1的f字段使用json:"foo"标签定义了json的key为foo，T2的f字段使用json:"bar"标签定义了json的key为bar。
//在main函数中，我们创建了一个T1类型的值t1，并打印它的值。然后，我们将t1的值转换为T2类型，并打印转换后的值。
//由于T1和T2的f字段都定义了json标签，所以转换后的值会使用json标签定义的key。
//因此，t2的值为{10}，而不是{20}。
