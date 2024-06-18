package main

import "fmt"

//在Go语言中，interface{}是一个空接口类型，它可以存储任何类型的值。由于它不声明任何方法，所以任何类型都满足interface{}类型。以下是一些关于interface{}的基本语法用法：

//直接赋值
type MyInterface interface{}

type MyStruct struct {
	name string
}

func (m MyStruct) String() interface{} {
	return m.name
}

func main() {
	//赋值list
	var myVar []interface{} //通用容器：由于interface{}可以存储任何类型，它经常用作通用的容器或切片，用于存储不同类型的值
	myVar = append(myVar, 1)
	myVar = append(myVar, "hello")
	myVar = append(myVar, true)
	myVar = append(myVar, []int{1, 2, 3})
	myVar = append(myVar, map[string]string{"key1": "value1", "key2": "value2"})
	myVar = append(myVar, func(x, y int) int { return x + y })
	fmt.Println(myVar)

	//赋值ma

	myMap := make(map[string]interface{})
	myMap["key1"] = 1
	myMap["key2"] = "hello"
	myMap["key3"] = true
	myMap["key4"] = []int{1, 2, 3}
	myMap["key5"] = map[string]string{"key1": "value1", "key2": "value2"}
	myMap["key6"] = func(x, y int) int { return x + y }
	fmt.Println(myMap)

	//赋值struct
	type myStruct struct {
		key1 int
		key2 string
		key3 bool
		key4 []int
		key5 map[string]string
		key6 func(x, y int) int
	}
	var myStructVar myStruct
	myStructVar.key1 = 1
	myStructVar.key2 = "hello"
	myStructVar.key3 = true
	myStructVar.key4 = []int{1, 2, 3}
	myStructVar.key5 = map[string]string{"key1": "value1", "key2": "value2"}
	myStructVar.key6 = func(x, y int) int { return x + y }
	fmt.Println(myStructVar)

}
