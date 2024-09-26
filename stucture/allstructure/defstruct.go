package allstructure

import "fmt"

//define struct

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

//结构体方法
func (p Person) Greeting() string {
	return "Hello, " + p.Name
}

func NewPerson() {
	// Initialize a new person struct instance
	p := &Person{"Alice", 30, "123 Wonderland"}
	fmt.Println(p.Greeting())
	fmt.Println(p)
	//也可以用字段名进行初始化
	q := &Person{Name: "Bob", Age: 25, Address: "456 Evergreen"}
	fmt.Println(q)
	fmt.Println(q.Address)
	//使用结构体字段
	p.Name = "Charlie"
	fmt.Println(p)
	//结构体指针
	var r *Person = &Person{"David", 35, "789 Ivy"}
	fmt.Println(r)
	//修改结构体指针
	r.Age = 40
	fmt.Println(r)
	// 结构体指针和值
	fmt.Println(p.Name)
	fmt.Println((*r).Name)

}
