package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type User struct {
	Name   string   `xml:"name"`
	Age    int      `xml:"age"`
	Email  string   `xml:"email,omitempty"`
	Phones []string `xml:"phone,omitempty"`
}

func main() {
	user := User{
		Name:   "John Doe",
		Age:    30,
		Email:  "",
		Phones: []string{"123-456-7890", "098-765-4321"},
	}

	// 序列化结构体为 XML
	xmlData, err := xml.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshalling XML: %v", err)
		return
	}

	// 输出 XML 数据
	fmt.Println(strings.Split(string(xmlData), "\n"))
	fmt.Println(string(xmlData))
	// 输出可能类似于：
	// <?xml version="1.0" encoding="UTF-8"?>
	// <User>
	// 	<name>John Doe</name>
	// 	<age>30</age>
	// 	<phone>123-456-7890</phone>
	// 	<phone>098-765-4321</phone>
	// </User>
}
