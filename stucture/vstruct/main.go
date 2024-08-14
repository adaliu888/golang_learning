// 把数据写入结构体
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Country string `json:"Country"`
}

func main() {

	var person Person
	//创建结构体实例
	log.Println(person)
	personwithdetail := Person{
		Name:    "alice",
		Age:     25,
		Country: "abc",
	}
	fmt.Println(personwithdetail)
	//直接通过（.）进行赋值
	person.Name = "bob"
	fmt.Println(person.Name)
	//声明的同时赋值
	p1 := Person{
		Name:    "hoh",
		Age:     12,
		Country: "jj",
	}
	fmt.Println(p1)

	//使用 new 函数创建一个结构体指针，然后通过指针间接访问和赋值字段。
	p2 := new(Person)
	p2.Country = "heol"
	fmt.Println(p2.Country)
	//函数体
	p3 := CreatePerson("jim", 33)
	fmt.Println(p3)

	//构造函数
	p4 := NewPerson("ah", 14, "bb")
	log.Fatal(p4)

	//使用 Map 来动态赋值:
	//如果你有一个字段名和值的映射，可以使用反射来动态地赋值
	personvalues := map[string]interface{}{
		"Name":    "map",
		"Age":     12,
		"Country": "zz",
	}
	for fieldName, value := range personvalues {
		fmt.Println(fieldName, value)

	}
	//
	jsonStr := `{"Name": "Grace", "Age": 22, "Country": "Hopperland"}`
	var p5 Person
	err := json.Unmarshal([]byte(jsonStr), &p5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p5.Name, p5.Age, p5.Country)

}

func CreatePerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}

}

// 使用构造函数
// 定义一个函数，该函数接收参数并返回一个初始化好的状态的实例
func NewPerson(name string, age int, country string) *Person {
	return &Person{
		Name:    name,
		Age:     age,
		Country: country,
	}
}
