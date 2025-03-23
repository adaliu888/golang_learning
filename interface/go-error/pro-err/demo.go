package main

import (
	"fmt"
)

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}
	return x / y, nil
}

func add(x, y int) (int, error) {
	if x+y < 0 {
		return x + y, fmt.Errorf("overflow")
	}

	return x + y, nil
}

func main() {
	//检查错误
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	}

	result1, err := add(10, -30)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result1)
	fmt.Println(result)
}
