package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := make(chan string)
	fmt.Println(v)

	r := reflect.TypeOf(v).Kind()
	fmt.Println(r)

	r1 := reflect.ValueOf(v)
	fmt.Println(r1)
}
