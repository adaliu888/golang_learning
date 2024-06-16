package main

import "fmt"

func main() {

	fmt.Println("hello")

	p0 := new(int)   //create a type , the type is int
	fmt.Println(p0)  //print the address of the int
	fmt.Println(*p0) //read the value of address p0

	x := p0
	fmt.Println(&x) // point to the type of address
	fmt.Println(x)  // point to the type of address
	fmt.Println(*x) //read the value of address

	p1, p2 := &x, &x
	fmt.Println(p1 == p2) //true
	//fmt.Println(p0 == p1)

	p3 := *&p0            //p3 := p0 *int
	fmt.Println(p3 == p0) //true

	var p4 *int

	fmt.Println(p4)
	fmt.Println(&p4)
	fmt.Println(*&p4) // point to the type of address
	p5 := &p4

	fmt.Println(*p5)
	fmt.Println(p5) //print the address of the int

}
