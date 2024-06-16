package main

import (
	"fmt"
	"sync"
)

func sum(numbers []int, c chan int) {
	toatal := 0
	for _, number := range numbers {
		toatal += number
	}
	c <- toatal //send total to c

}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup

	c := make(chan int)
	go sum(numbers[:len(numbers)/2], c)
	go sum(numbers[len(numbers)/2:], c)

	x := <-c
	y := <-c
	wg.Wait()
	fmt.Println(x, y, x+y)

}
