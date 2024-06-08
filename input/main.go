package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// create func input
func getInput(promote string, r *bufio.Reader) (string, error) {
	fmt.Println(promote)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

// create new bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	//fmt.Print("create a new bill name:")
	//name, _ := reader.ReadString('\n')
	//name = strings.TrimSpace(name)
	name, _ := getInput("create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("create the bill: ", b.name)
	return b
}

// create func promoteoperate
func promoteOperate(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("a - add a item,s - save bill, t - add tip ", reader)
	fmt.Println(opt)
	switch opt {
	case "a": //input k,v in items
		name, _ := getInput("item name:", reader)
		price, _ := getInput("item price:", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promoteOperate(b)
		}
		b.addItem(name, p)
		fmt.Println("the item", name, price)
		promoteOperate(b)
	case "s": //switch s ,call function save() to save bill file
		b.save()
		fmt.Println("save bills file", b.name)
	case "t":
		tip, _ := getInput("enter tip amount ($):", reader)
		p, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promoteOperate(b)
		}
		b.updateTip(p)
		fmt.Println("the item", tip)
		promoteOperate(b)
	default:
		fmt.Println("you input a vild number")
		promoteOperate(b)

	}
}

func main() {
	mybill := createBill()
	promoteOperate(mybill)
	fmt.Println(mybill)

}
