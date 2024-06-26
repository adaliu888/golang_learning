package main

import "fmt"

type myinput interface { //define a interface myinput,ss is a method
	ss(s string) string //method ss
}

type myFunc func(s string) string // var a myFunc is a func and parameter is (s string) and return a string

func (*myFunc) ss(s string) string {
	return fmt.Sprintf("%s\n", s)
}

var _ myinput = (*myFunc)(nil)

func main() {
	var mf myFunc
	fmt.Println(mf.ss("hello"))
}
