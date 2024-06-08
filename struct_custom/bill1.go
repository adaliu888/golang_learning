package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newbill(name string) bill {
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
