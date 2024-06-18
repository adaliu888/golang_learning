package main

import (
	"fmt"
	"os"
)

type Person struct {
	name string
	age  int
}

// 有传参和返回值
func InputName(name *Person) string {
	return name.name
}

// 有传参和返回值
func InputAge(age *Person) int {
	return age.age
}

// 没有传参，有返回值
func OutputName() string {
	return "John"
}

// 返回0或者多个返回值
func Divide(x float32, y float32) (float32, float32) {
	if x/y == 0 {
		return 0, y
	} else {
		return x, y
	}

}
func main() {
	p := Person{ //字面量
		name: "John",
		age:  30,
	}
	fmt.Println(p)
	fmt.Println(InputName(&p))
	fmt.Println(InputAge(&p))
	fmt.Println(p.name)
	fmt.Println(p.age)
	//无传参，无返回值
	fmt.Println(OutputName())

	//返回多个返回值
	a, b := Divide(0, 5)
	fmt.Println(a, b)
	//匿名函数
	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
	//闭包  函数使用之外的变量
	c := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}(1)
	fmt.Println(c(2))
	fmt.Println(factorial(5))
	//错误处理
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println(f)

	// 无错误处理
	f, _ = os.Open("a.txt")
	defer f.Close()
	fmt.Println(f)

	//高阶函数 go语言的函数是一等公民，可以作为参数传递到其他函数中，也可以作为返回值返回到其他函数中

	//func LoopWord(times int,f func()) {
	//for i := 0; i < times; i++ {
	//	f()
	//}
	//递归函数

}

// 递归函数
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
