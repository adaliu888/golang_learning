package main

import (
	"fmt"
)

func outer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in outer:", r)
		}
	}()

	fmt.Println("Executing outer function")
	middle()
}

func middle() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in middle:", r)
		}
	}()

	fmt.Println("Executing middle function")
	inner()
}

func inner() {
	fmt.Println("Executing inner function")
	panic("A panic occurred")
}

func main() {
	outer()
	fmt.Println("After calling outer")
	inner()

	fmt.Println("After calling inner")
	middle()
	fmt.Println("After calling middle")
}
