package main

import (
	"ada/integers"
	"fmt"
)

var five = 5

func main() {

	{
		var integer = 2
		fmt.Println(integer)
	}
	var integer = 1

	fmt.Println("hello world")
	fmt.Println(integer)
	fmt.Println(five)
	fmt.Println(integers.Three)

}
func nonmain() {
	fmt.Println(five)

}
