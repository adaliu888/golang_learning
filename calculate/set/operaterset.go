package operaterset

import (
	"fmt"
	"reflect"
)

func OperaterSet() {
	//创建一个fruits map
	var fruits map[string]int
	fruits = make(map[string]int)
	fruits["apple"] = 100
	fruits["banana"] = 200
	fruits["orange"] = 300
	fruits["mango"] = 400
	fruits["grape"] = 500
	fmt.Println(fruits)

	//判断key是否存在
	if _, ok := fruits["apple"]; ok {
		fmt.Println("apple key exists")
	}
	//get values from key
	fmt.Println(fruits["apple"])

	//delete key
	delete(fruits, "apple")
	fmt.Println(fruits)

	//check if key exists
	if _, ok := fruits["apple"]; !ok {
		fmt.Println("apple key  not exists")
	}
	// add element
	fruits["kiwi"] = 600
	fmt.Println(fruits)
	// update element
	fruits["kiwi"] = 700
	fmt.Println(fruits)

	//遍历
	for k, v := range fruits {
		fmt.Println(k, v)
	}
	//predicate true
	r := reflect.TypeOf(fruits)

	if r.Kind() == reflect.Map {
		fmt.Println("fruits is a map")
	}
	//get length
	fmt.Println(len(fruits))
	//check if map is empty
	if len(fruits) == 0 {
		fmt.Println("map is empty")
	} else {
		fmt.Println("map is not empty")
	}
	//search key
	fmt.Println(fruits["apple"])

	//clear map
	fruits = make(map[string]int)
	fmt.Println(fruits)

	//create list
	var arr []string = []string{"apple", "banana", "orange", "mango", "grape"}

	//add arr
	arr = append(arr, "kiwi")
	fmt.Println(arr)
	//delete value
	arr = arr[:len(arr)-2] //切片删掉最后两个
	fmt.Println(arr)
	//add arr
	arr = append(arr, "kiwi", "grape")
	fmt.Println(arr)
	//遍历
	for _, v := range arr {
		fmt.Println(v)
	}
	//get length
	fmt.Println(len(arr))
	//check if list is empty
	if len(arr) == 0 {
		fmt.Println("list is empty")
	} else {
		fmt.Println("list is not empty")
	}
	//search value
	if arr[0] == "apple" {
		fmt.Println("apple")
	}

	//clear list
	arr = []string{}
	fmt.Println(arr)

}
