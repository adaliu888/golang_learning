package main

import (
	"fmt"
	"strings"
)

func main() {
	//使用常量
	const (
		Months      = `Jan\0Feb\0Mar\0Apr`
		MonthsCount = 4
	)
	//var monthSlice []string = make([]string, MonthsCount)
	monthSlice := strings.Split(Months, `\0`)
	//copy(monthSlice, []string(Months))

	fmt.Printf("%+v\n", monthSlice) // Jan\0Feb\0Mar\0Apr\n(monthSlice)
}
