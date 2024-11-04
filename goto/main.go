package main

import "fmt"

func main() {
	i := 0
Loop:
	for i < 3 {
		fmt.Println("Inside loop:", i)
		i++
		if i == 2 {
			goto Loop
		}
	}
	fmt.Println("Loop exited.")
}
