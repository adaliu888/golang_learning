package main

//define the error with the struct

import (
	"fmt"
)

type DivideError struct {
	Dividee float64
	divider float64
}

// 实现 `error` 接口
func (de *DivideError) Error() string {  //实现一个error接口
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.divider)
}

func Divide(varDividee float64, varDivider float64) (result float64, err error) {
	if varDivider == 0 {
		return 0, &DivideError{varDividee, varDivider}

	}
	return varDividee / varDivider, err

}

func main() {
	fmt.Println(Divide(100.0, 10.0))
	fmt.Println(Divide(100.0, 0.0))
}
