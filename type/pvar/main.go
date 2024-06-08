package pvar

import "fmt"

func double(x *int64) {
	*x += *x
	x = nil //此行仅解释目的
}
func main() {
	var a int = 3
	double(&a) /
	fmt.Println(a)
	p := &a
	fmt.Println(a, p == nil)
}
