package main

import "fmt"

type T struct {
	f string `json:"one two three"`
}

func main() {
	t := T{}
	t.f = "hello world"
	fmt.Printf("%#v\n", t)
}
