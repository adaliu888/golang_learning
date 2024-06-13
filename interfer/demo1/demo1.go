package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() { //returns varible is every type
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone) //type NokiaPhone struct(),func (nokiaPhone,nokiaPhone) call()
	phone.call()            //called call()

	phone = new(IPhone) //type Iphone struct{}, func (iphone IPhone) call()
	phone.call()
}
