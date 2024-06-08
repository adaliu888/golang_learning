package main // every program must have a main package to be compiled

import (
	"fmt"
	"strings" //golang standard libray
) //fmt is a standard library package that contains functions such as print() and output()

func main() {
	str_aa := "please rememenber me " //this is about usage of the standard library
	fmt.Println(str_aa)
	fmt.Println(strings.Contains(str_aa, "please"))
	fmt.Println(strings.ReplaceAll(str_aa, "please", "do you"))
	fmt.Println(strings.ToUpper(str_aa))
	fmt.Println(strings.Count(str_aa, "m"))

	//for loop
	names := []string{"change", "world", "golang", "kess"} //this is about for loop.
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}
	for index, value := range names { //range loop
		if index == 1 {
			fmt.Println(names[index], names)
		}
		fmt.Println("print the index and values:", index, value)
	}
	//fmt.Println("Hello, World!") /*.this is a comment.use it to make your code more readable...*/
	//--------------------------------------------------------------------------
	//int type
	//var a int = 10
	//var b int = 20
	//var c int = a + b
	//var s string = "hello, " + "world"
	//fmt.Println(c, s)
	//--------------------------------------------------------------------------
	//string type
	//var nameone string = "goodbye "
	//var nametwo string = "world"
	//var namethree = "hello, " + "world" // this is an expression,:= and = is an assignment operator
	//fmt.Println(nameone, nametwo, namethree)
	//nameone = "hello, " + "world"
	//fmt.Println(nameone, nametwo, namethree)

	//namefour := "hello,yun"
	//fmt.Println(namefour)

	//bits memory
	//var onenum uint8 = 255
	//var twonum uint16 = 65535
	//var threenum uint32 = 4294967295
	//var fournum uint64 = 18446744073709551615
	//fmt.Println(onenum, twonum, threenum, fournum)

	//float memory
	//var f1 float32 = 3.14
	//var f2 float64 = 3.141592653589793
	//f3 := 3.141592653589793

	//fmt.Println(f1, f2, f3)

	//bool memory
	//var b1 bool = true
	//var b2 bool = false
	//b3 := true
	//fmt.Println(b1, b2, b3)

	//char memory
	//var c1 byte = 'a'
	//var c2 byte = 'b'
	//c3 := 'c'
	//fmt.Println(c1, c2, c3)

	//string memory
	//var s1 string = "hello, " + "world"
	//s2 := "hello, " + "world"
	//fmt.Println(s1, s2)

	//array memory
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	arr[0] = 0

	fmt.Println(arr)

	//slice memory
	var sl []int = []int{1, 2, 3, 4, 5}
	fmt.Println(sl)

	//map memory
	var mp map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	var mp1 = make(map[string]int, 10)
	mp1["a"] = 1
	mp1["b"] = 2
	mp1["c"] = 3
	mp1["d"] = 4
	mp1["e"] = 5

	fmt.Println(mp, mp1)
	mp1["a"] = 0 //modify map value
	mp1["f"] = 6 //add map value

	fmt.Println(mp, mp1)

	//struct memory as type system
	type Person struct {
		Name string
		Age  int
		Sex  string
	}
	var p Person = Person{}
	p.Name = "yun"
	p.Age = 20
	p.Sex = "male"
	fmt.Println(p)

	//interface memory
	var i interface{}
	i = 1
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)

	i = true
	fmt.Println(i)

	//pointer memory
	var p1 *int
	p1 = &arr[0]
	fmt.Println(p1)

	var p2 *string
	p2 = &p.Name
	fmt.Println(p2)
	b1 := true
	var p3 *bool
	p3 = &b1
	fmt.Println(p3)

	//print memory
	fmt.Printf("Hello, World!\n")

	//for memory
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//if memory
	if true {
		fmt.Println("hello, world")
	}

	//switch memory
	switch 2 {
	case 1:
		fmt.Println("hello, world")
	case 2:
		fmt.Println("hello, yun")
	default:
		fmt.Println("hello, ada")
	}

	//defer memory
	defer fmt.Println("hello, world")
	defer fmt.Println("hello, yun")
	defer fmt.Println("hello, ada")

	//go memory
	go fmt.Println("hello, world")
	go fmt.Println("hello, yun")
	go fmt.Println("hello, ada")

	name := "yun"
	age := 20
	gender := "male"
	fmt.Printf("name:%s, age:%d, gender:%s\n", name, age, gender)

}
