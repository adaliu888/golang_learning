package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	fmt.Println(b)
	return b
}

// format the bill
func (b bill) format() string {
	fs := "bill breakdown:"
	var total float64 = 0

	//list item
	for k, v := range b.items {
		fmt.Printf("%v ...$%v", k+":", v)
		total += v
	}
	//total
	fs += fmt.Sprintf("%v ...$%0.2f", "total:", total)
	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add a item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save file
func (b *bill) save() { //do not return premar
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("save bills file")

}
