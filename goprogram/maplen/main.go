package main

import (
	"fmt"
)

func main() {
	//创建一个map
	Mymap := make(map[string]int)

	//添加键值对
	Mymap["one"] = 1
	Mymap["two"] = 2
	Mymap["three"] = 3

	//判断map的长度为0

	if len(Mymap) == 0 {
		fmt.Println("map is empty")
	} else {
		fmt.Println("map length is ", len(Mymap))

	}

}
