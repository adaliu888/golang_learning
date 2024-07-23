package main

import (
	"fmt"
)

// define long string println
func longstr(s string, n int) string {
	for i := 0; i < n; i++ {
		return fmt.Sprintf("%s\n", s)

	}
	return "" // unreachable code, but required for function signature.
}

func main() {

	fmt.Println(longstr("Hello, World!", 5))

}
