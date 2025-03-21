package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func NewBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 2.99},
		tip:   0,
	}
	return b
}

func (b bill) format() string { //b bill as format() is description
	fs := "bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%s ...$%.2f \n", k+":", v)
		total += v
	}

	// add tip to total
	total += b.tip

	// total with tip
	fs += fmt.Sprintf("Total (including tip) ...$%.2f \n", total)

	return fs
}

//main 函数创建了一个 bill 实例，设置了小费，并打印了格式化的账单
func main() {
	mybill := NewBill("mario is bill")

	fmt.Println(mybill.format())

}
