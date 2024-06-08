package main

import "fmt"

func main() {
	//create new map
	menu := map[string]float64{
		"hamberger": 34.5,
		"icecream":  6.3,
		"cocolo":    7.5,
		"freshfire": 12,
	}
	fmt.Println(menu)              //print map
	fmt.Println(menu["hamberger"]) //print map menu
	//print maps menu k,v
	for k, v := range menu {
		fmt.Println(k, v)
	}
	//create new map phonebook,and print map and fliter keys
	phonebook := map[int]string{
		2345678: "mario",
		9876543: "harli",
		3456879: "lsihis",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[2345678])

	phonebook[2345678] = "bowers"
	fmt.Println(phonebook)

}
