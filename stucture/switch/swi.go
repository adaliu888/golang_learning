package swi

import (
	"fmt"
)

func Print[T any](elementS []T) {
	for _, element := range elementS {
		fmt.Printf("%v ", element)
	}

}

func PrintMain() {
	Print([]int{1, 2, 3})
	Print([]string{"a", "b", "c"})
}
