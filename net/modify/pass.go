package main

import "fmt"

func updateName(x string) string { //define function updatename
	x = "wedge" //update data is wedge
	return x
}

//define function to change value of the map of y
func updatemenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	//group a types-> strings,ints,bools,floats,arrays,structs
	name := "tifa"
	name = updateName(name) //call function updateName
	fmt.Println(name)

	menu := map[string]float64{
		"coffee": 7.66,
		"juice":  1.99,
	}
	updatemenu(menu)
	fmt.Println("memory address of name is ", menu)
	m := &name
	fmt.Println("value at memory address", *m)
	fmt.Println("memory address", m)

	fmt.Println(menu)

}
