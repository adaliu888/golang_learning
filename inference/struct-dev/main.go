package main

import "fmt"

type FirstInterface interface {
	Firstone()
}

type SecondInterface interface {
	FirstInterface //interface嵌入interface,是指在接口中嵌入接口。
	Secondone()
	Thirdone()
}

type MyStruct struct {
	//匿名匿名嵌入是指在结构体中不命名嵌入的字段，只指定类型。这种嵌入方式可以自动实现嵌入类型的接口方法。
	//特点：
	//结构体自动实现嵌入类型的所有接口方法。
	//嵌入的字段不能直接访问，但可以通过类型断言或反射间接访问。
	//嵌入相同类型的所有实例共享同一个实例。
	SecondInterface
}

func (m MyStruct) Firstone() {
	fmt.Println("Firstone")
}

func (m MyStruct) Secondone() {
	fmt.Println("Secondone")
}

func (m MyStruct) Thirdone() {
	fmt.Println("Thirdone")
}

func main() {
	var SecondInterface SecondInterface

	fmt.Println(SecondInterface)

	SecondInterface.Secondone()

}
