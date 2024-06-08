package main

import (
	"fmt"
)

func main() {
	//fmt.Println("this is another golang main")
	//for loop and if condition,thought if continue to control the excute flow.
	names := []string{"mango", "banana", "apple", "pear"}
	for index, value := range names {
		//fmt.Println(index, value)
		if index == 1 {
			fmt.Println("print the index1", index)
			continue
		} else if index == 2 {
			fmt.Println("print the index2", index)
		}
		if index > 2 {
			fmt.Println("the program stop excute")
			break
		}
		fmt.Println(index, value)

	}

}
