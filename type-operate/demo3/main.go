package main

import "fmt"

//long type

func longarr(args []string) {
	for _, v := range args {
		fmt.Println(v)
	}
}
func longstr(s string, n int) string {
	for i := 0; i < n; i++ {
		fmt.Sprintf("%s\n", s)
	}
	return "" // unreachable code, but required for function signature.
}

//struct type
type Longstruct struct {
	s string
	n int
}

func main() {
	longarr([]string{"123", "456", "789"}) //describe a long array
	longstr("Hello, World!", 5)
	Longstruct := &Longstruct{"name", 5}
	longstr(Longstruct.s, Longstruct.n) //describe a long struct

}
