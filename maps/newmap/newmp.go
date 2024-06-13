package main

import "fmt"

var map1 = map[string]string{
	"name": "johon",
	"sex":  "famel",
}

func main() {
	fmt.Println("hello world")
	fmt.Println(map1)
	fmt.Println(map1["name"])
	fmt.Println(map1["sex"])
	fmt.Println(len(map1))
	for _, v := range map1 {
		fmt.Println(v)
	}
	delete(map1, "name")
	fmt.Println(map1)

}
