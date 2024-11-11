package main

import (
	"fmt"
)

// 如何判断 map 中是否包含某个 key
func main() {
	Mymap := map[string]string{
		"name": "John",
		"age":  "25",
		"city": "New York",
	}

	key := "name"

	if _, ok := Mymap[key]; ok {
		fmt.Printf("The key %s is in the map\n", key)
	} else {
		fmt.Printf("The key %s is not in the map\n", key)
	}

	// 如何删除 map 中的 key-value 项
	delete(Mymap, key)

	fmt.Println("After deleting the key-value pair:", Mymap)
}
