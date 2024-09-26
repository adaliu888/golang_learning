package allstructure

import (
	"fmt"
)

type Details struct {
	Age   int
	Email string
}

type User struct {
	Name string
	Details
}

func WriteDetails() {
	u := User{Name: "Alice", Details: Details{Age: 30, Email: "alice@example.com"}}
	fmt.Println(u.Name, u.Details, u.Email) //匿名嵌套
	//fmt.Println(u.Details.Age, u.Details.Email) //有名嵌套
}
