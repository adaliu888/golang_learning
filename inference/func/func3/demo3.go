package main

import "fmt"

type MyInterface interface { //定义一个接口,MyInterface是空接口
} //定义一个空接口

type MyStruct struct {
	Impress MyInterface //
	name    string
}

func (m *MyStruct) String() string {
	return m.name
}

//type Impress interface { //定义一个接口
//	PS() string
//}

func (m *MyStruct) PS() string {
	return m.name
}

func main() {
	Mystruct := &MyStruct{
		name: "hello",
	}
	fmt.Println(Mystruct)
	fmt.Println(Mystruct.PS())
}
