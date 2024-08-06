package main

import (
	"fmt"
	"reflect"
)

func main() {

	var m map[string]int

	if m == nil {
		fmt.Println("m is nil")
	}
	result := reflect.ValueOf(m)
	fmt.Println(result)

	p := &m
	fmt.Println(*p)

	fmt.Println(&m)

}
