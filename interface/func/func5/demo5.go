package main

import (
	"fmt"
	"log"
)

type PersonAction interface {
	speak() string
	run() string
	comunicate() string
}
type Person struct {
	name string
	age  int
}

func Name(p *Person) string {
	return p.name
}

func Age(p *Person) int {
	return p.age
}

func (p *Person) Speak() string {
	return "I am " + p.name
}

func (p *Person) Run() string {
	return "I am running" + p.name
}

func (p *Person) Communicate() string {
	return "I am talking" + p.name
}

func main() {

	p := Person{
		name: "John",
		age:  30,
	}

	log.Println(p.Speak())

	fmt.Println(p.Run())

	fmt.Println(p.Communicate())
}
