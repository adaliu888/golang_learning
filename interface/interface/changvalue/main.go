package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value interface{}

	// 动态设置 value 的类型和值
	value = 42
	fmt.Println("value:", value)

	// 使用反射来检查 value 的类型和值
	v := reflect.ValueOf(value)
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Interface())

	// 动态地改变 value 的类型
	value = "Hello, world!"
	fmt.Println("new value:", value)

	// 再次使用反射来检查
	v = reflect.ValueOf(value)
	fmt.Println("new type:", v.Type())
	fmt.Println("new value:", v.String())
}
