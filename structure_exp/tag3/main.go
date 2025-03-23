package main

import (
	"fmt"
	"reflect"
)

// 通过空格符来分割键值 —key1:"value1" key2:"value2" key3:"value3"。如果Tags格式没问题的话，我们可以通过Lookup或者Get来获取键值对的值。
// Lookup回传两个值 —对应的值和是否找到
type T struct {
	f string `one:"1" two:"2"blank:""` //tag keys with values is null ,not print
}

func main() {
	t := reflect.TypeOf(T{})
	f, _ := t.FieldByName("f")
	fmt.Println(f.Tag) // one:"1" two:"2"blank:""
	v, ok := f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
	v, ok = f.Tag.Lookup("blank")
	fmt.Printf("%s, %t\n", v, ok) // , true
	v, ok = f.Tag.Lookup("five")
	fmt.Printf("%s, %t\n", v, ok) // , false
}
