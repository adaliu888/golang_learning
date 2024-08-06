package main

import (
	"fmt"
	"unsafe"

)

func main() {
	const (
		Months      = "Jan\0Feb\0Mar\0Apr"
		MonthsCount = 4
	)
	var monthSlice []string = make([]string, MonthsCount)
	copy(monthSlice, []byte(Months))

	fmt.Println(monthSlice)
}
