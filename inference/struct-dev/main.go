package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	name() string
}

type MyStruct struct {
	name string
}

func NewMyStruct(name string) *MyStruct {
	return &MyStruct{name: name}
}

func main() {

	myStruct := NewMyStruct("hello")

	fmt.Println(reflect.TypeOf(myStruct).Elem())
	fmt.Println(reflect.TypeOf(myStruct).Elem().Name())

	fmt.Println(myStruct)

	fmt.Println(reflect.TypeOf(43))
	fmt.Println(reflect.TypeOf(43).Kind())
	fmt.Println(reflect.TypeOf(43).Name())

	fmt.Println(reflect.TypeOf("hello").Kind())

}
