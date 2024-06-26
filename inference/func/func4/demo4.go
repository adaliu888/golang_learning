package main

import "fmt"

type MyInterface interface {
	// 空接口，不包含任何方法
}

// MyStruct 嵌入一个实现了 MyInterface 的具体类型的实例
type MyStruct struct {
	implementor // 假设 implementor 是实现了 MyInterface 的类型
	name        string
}

// 假设下面的类型实现了 MyInterface
type implementor struct{}

func (i implementor) Method() {
	// 实现 MyInterface 接口的方法
}

// 现在创建一个 MyStruct 的实例，并嵌入实现了 MyInterface 的 Implementor
func main() {
	myImplementor := implementor{}
	myStruct := MyStruct{
		implementor: myImplementor,
		name:        "MyStruct instance",
	}

	fmt.Println(myStruct)

}
