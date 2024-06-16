package main

import "fmt"

type Person struct {
	name    string
	age     int
	Address //非匿名嵌套，作为嵌套字段
}

type Address struct {
	city string
	zip  int
}

type Employee struct {
	Name    string
	Address // 匿名嵌套，可以直接访问 Street 和 City
}

func main() {
	person := Person{
		name: "John",
		age:  30,
		Address: Address{
			city: "Beijing",
			zip:  100000,
		},
	}
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.Address)
	fmt.Print(person.Address.city) //非匿名嵌套，可以直接访问 city

	Employee := Employee{
		Name: "John",
		Address: Address{
			city: "Beijing",
			zip:  100000,
		},
	}
	fmt.Println(Employee.Name)

	fmt.Println(Employee.Address.city)
	fmt.Println(Employee.zip) //匿名嵌套，可以直接访问 zip 和 City
}
