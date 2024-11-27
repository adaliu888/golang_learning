package main

import (
	"fmt"
)

func main() {
	str := "Hello, 世界"
	runes := []rune(str) //print char though using rune conversion
	for i, r := range runes {
		fmt.Printf("Position %d: %c\n", i, r)
	}
	// Output:
	// Position 0: H
	// Position 1: e
	// Position 2: l
	// Position 3: l
	// Position 4: o
	// Position 5: ,
	// Position 6: 世
	// Position 7: 界

	bytes := []byte(str)
	for i, b := range bytes {
		fmt.Printf("Position %d: %c\n", i, b)
	}
	// Output:

}
