package main

import (
	"fmt"
	"math"
)

// define and call function
func saygreeting(n string) {
	fmt.Printf("say good morning %v \n", n)
}
func saybye(n string) {
	fmt.Printf("say good bye %v \n", n)
}
func namecircle(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}

}
func circlearea(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	fmt.Printf("this is main func \n")
	saybye("gile")
	saygreeting("mile")
	namecircle([]string{"heloo", "holo"}, saygreeting)
	circlearea(7.0)
	a1 := circlearea(10.5)
	a2 := circlearea(9.5)
	fmt.Println(a1, a2)

}
