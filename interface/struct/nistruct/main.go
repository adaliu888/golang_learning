package main

import "fmt"

type InnerStruct struct {
	InnerField string
}

func (is *InnerStruct) InnerMethod() {
	fmt.Println("InnerMethod called")
}

type OuterStruct struct {
	InnerStruct // 匿名嵌套 InnerStruct
}

func main() {
	outer := OuterStruct{}
	outer.InnerField = "Hello, Inner World!"

	// 直接访问匿名嵌套的字段
	fmt.Println(outer.InnerField)

	// 通过匿名嵌套调用 InnerStruct 的方法
	outer.InnerStruct.InnerMethod()
}
