package main

import (
	"fmt"
	"reflect"
)

type T struct {
	f1     string `f one`
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
}

type S struct {
	s1 string ` `
	s2 string
}

func main() {
	t := reflect.TypeOf(T{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag) // f one
	f2, _ := t.FieldByName("f2")
	fmt.Println(f2.Tag) // f two
	f3, _ := t.FieldByName("f3")
	fmt.Println(f3.Tag) // f three
	f4, _ := t.FieldByName("f4")
	fmt.Println(f4.Tag) // f four and five
	f5, _ := t.FieldByName("f5")
	fmt.Println(f5.Tag) // f four and five

	//
	s := reflect.TypeOf(S{})
	f1, _ = s.FieldByName("s1")
	fmt.Println(f1.Tag) // null
	f2, _ = s.FieldByName("s2")
	fmt.Println(f2.Tag) // null

}
