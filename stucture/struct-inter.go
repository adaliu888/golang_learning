// 接口也可以嵌套其他接口，允许更复杂的行为组合。
package main

import "fmt"

// 定义基本接口
type Greeter interface {
	Greet()
}

// 定义工作接口，嵌套 Greeter 接口
type Worker interface {
	Work()
	Greeter // 嵌套接口
}

// 定义结构体
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s.\n", p.Name)
}

func (p Person) Work() {
	fmt.Printf("%s is working.\n", p.Name)
}

func main() {
	var w Worker = Person{Name: "Alice"}
	w.Greet() // 调用嵌套接口的方法
	w.Work()  // 调用 Worker 接口的方法
}
