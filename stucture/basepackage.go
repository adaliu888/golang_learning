package main //go language management code thought package to manage

import (
	"fmt" //i/o conversion
	"math/rand"
	"strings"
	"time"
)

var bob = titleName("bob")
var smith = titleName("smith")

// init function will be executed before main function
func init() {
	fmt.Println("hi", bob)

}

func init() {
	fmt.Println("hello", smith)
}
func titleName(s string) string {
	s = strings.Title(s)
	return s
}

func main() { //define main to execute the command
	fmt.Println("go has", 25, "keyboards: ")
	//print the keyboards from the command line
	fmt.Println("next random number is : ", rand.Uint32()) //

	rand.Seed(time.Now().UnixNano())
	fmt.Println("next random number is : ", rand.Uint32())

	a, b := "go", "123"
	fmt.Printf("a == %v, b == %s \n", a, b)

	fmt.Println("bye bye")

}
