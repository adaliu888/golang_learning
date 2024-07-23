package main

import "fmt"

var map1 = map[string]string{
	"name": "johon",
	"sex":  "famel",
}

func main() {
	fmt.Println("hello world")
	fmt.Println(map1)         //println 会打印出 map[sex:famel name:johon]
	fmt.Println(map1["name"]) //println 会打印出 johon
	fmt.Println(map1["sex"])  //println 会打印出 famel
	fmt.Println(len(map1))    //println 会打印出 2
	for _, v := range map1 {  //println 会打印出 johon famel
		fmt.Println(v)
	}
	delete(map1, "name") //delete 会删除 name
	fmt.Println(map1)    //println 会打印出 map[sex:famel]
	map1["age"] = "20"   //添加 key-value
	map1["name"] = "alice"
	fmt.Println(map1) //println 会打印出 map[age:20 name:alice sex:famel]
} //add key-value
