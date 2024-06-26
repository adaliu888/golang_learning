package main

import "fmt"

// 函数调用自身

func factorial(x int) int {

	if x == 0 {
		return 1
	}
	resulte := x * factorial(x-1)
	fmt.Printf("%d", resulte)
	return resulte
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func input() int {
	return input()
}

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(factorial(100))
	fmt.Println(input())

}
