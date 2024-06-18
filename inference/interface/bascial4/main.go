//类型断言

package main

import "fmt"

func main() {

	var x interface{}
	//单返回值形式
	s := x.(int)
	fmt.Println(s)
	//双返回值形式
	x = 1

	if i, ok := x.(int); ok {
		fmt.Println(i)
	}

	x = "hello"
	if s, ok := x.(string); ok {
		fmt.Println(s)
	}

}
