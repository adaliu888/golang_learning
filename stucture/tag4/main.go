package main

import (
	"fmt"
)

type T1 struct {
	f int `json:"foo"`
}
type T2 struct {
	f int `json:"bar"`
}

func main() {
	t1 := T1{10}
	var t2 T2       // t2 as T2 type
	t2 = T2(t1)     //t1 as value for t2 of struct T2
	fmt.Println(t2) // {10}
}
