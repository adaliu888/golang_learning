package main

import "fmt"

//difine interface,接口通过关键字interface定义，可以包含一个或多个方法签名
type MyInterface interface {
	doWork() string
	getData() string
}
type MyInterface2 interface{} //define empty interface

//define struct

//实现接口,任何类型只要实现了接口中定义的所有方法，就被认为是实现了该接口。不需要显式声明：
type MyStruct struct {
	name string
}

func (m MyStruct) String() string {
	return m.name
}

func (m *MyStruct) doWork() string {
	return "do work"
}

func (m *MyStruct) getData() int {
	return 42
}

//在这个例子中，MyStruct实现了MyInterface接口

//使用接口值,接口值可以存储实现了接口的任何类型的值：
func main() {
	MyStruct := &MyStruct{
		name: "hello",
	}

	fmt.Println(MyStruct)
	MyInterface := MyStruct //把MyStruct赋值给MyInterface
	fmt.Println(&MyInterface)
	fmt.Println(MyInterface.doWork())
	fmt.Println(MyInterface.getData())
}
