package main

import "fmt"

type MyInterface interface{}

type MyStruct struct {
	name string
}

func (m *MyStruct) String() string {
	return m.name
}

func (m *MyStruct) doWork() string {
	return "do work"
}

// 由于MyInterface是空接口，MyStruct可以赋值给它
func main() {
	var myStructVar MyStruct
	myStructVar.name = "hello"
	fmt.Println(myStructVar) // 输出: hello

	// 将MyStruct赋值给MyInterface
	var myInterfaceVar MyInterface = myStructVar // MyInterface is type type
	fmt.Println(myInterfaceVar)                  // 输出: hello，调用了MyStruct的String()方法

	// 由于MyInterface是空接口，我们不能直接调用doWork方法
	// 但我们可以直接调用myStructVar.doWork()
	doWorkResult := myStructVar.doWork()
	fmt.Println(doWorkResult) // 输出: do work
}
