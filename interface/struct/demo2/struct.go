package main

import (
	"fmt"
)

type Driver struct {
	ch chan int
}

func (o *Driver) Constructor() {
	o.ch = make(chan int)

	o.ch <- 1

	go func() {
		a := <-o.ch
		fmt.Printf("a: %v\n", a)
	}()

}

type Person struct {
	Next *Person
}

func (o *Person) Constructor(next *Person) {
	o.Next = next
}

func doWork() {
	fmt.Println("do work")
}

type Employee struct {
	doWork func()
}

func (o *Employee) Constructor(doWork func()) {
	o.doWork = doWork
}

type Mystruct struct {
	field1 string
	filed2 string
}

func (o *Mystruct) Constructor(field1 string, filed2 string) { //use Constructor to struct varible fields,*Mystruct
	o.field1 = field1
	o.filed2 = filed2
}

// 接口
type Animal interface {
	string() string
}

type cat struct {
	Animal interface{}
}

func (c *cat) speak() string {
	return "meow"
}
func main() {
	//基本类型
	mystruct := &Mystruct{}
	fmt.Println(*mystruct)
	mystruct.Constructor("Hello", "World")
	mystruct.field1 = "Goodbye"
	fmt.Println(mystruct.field1)
	fmt.Println(mystruct.filed2)
	fmt.Println(*mystruct)
	//嵌套指针，指向相同类型的指针
	person := &Person{}
	fmt.Println(person)
	fmt.Println(*person)
	person.Constructor(person) //工厂模式
	fmt.Println(person.Next)
	fmt.Println(Person{})
	// 嵌套函数
	employee := &Employee{} //工厂模式
	employee.Constructor(doWork)
	employee.doWork()
	//嵌套接口
	newcat := &cat{}
	fmt.Println(newcat.speak())
	//嵌套管道
	newch := &Driver{}
	fmt.Println(newch.ch)

}
