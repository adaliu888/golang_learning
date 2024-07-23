package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic:", r)
		}
	}()
	panic("some panic")

}
