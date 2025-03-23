package main

import "fmt"

type NestedStruct struct {
	Field1 string
	Field2 int
}

type OuterStruct struct {
	NestedStruct
}

func (o *OuterStruct) Constructor(field1 string, field2 int) {
	o.NestedStruct.Field1 = field1
	o.NestedStruct.Field2 = field2
}

func main() {
	outer := &OuterStruct{}
	outer.Constructor("Hello", 42)

	fmt.Println(outer) // 输出: Hello

	outer.NestedStruct.Field1 = "Goodbye"

	fmt.Println(outer.NestedStruct.Field1) // 输出: Hello
	fmt.Println(outer.NestedStruct.Field2) // 输出: 42
}
