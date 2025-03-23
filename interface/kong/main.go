package main

import "fmt"

func PrintValue(v interface{}) {
	fmt.Println(v)
}

func main() {
	PrintValue(42)
	PrintValue("hello")
	PrintValue([]int{1, 2, 3})
}
