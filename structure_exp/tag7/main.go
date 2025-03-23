package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	// 假设这是从文件、网络或其他来源获取的 JSON 数据
	jsonData := []byte(`{"name": "John", "age": 30, "email": "KQY9C@example.com"}`)

	// 创建一个 User 结构体的指针
	var user *User

	// 使用 Unmarshal 方法解析 JSON 数据
	err := json.Unmarshal(jsonData, &user)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 打印反序列化后的数据
	fmt.Printf("User: %+v\n", *user)
}
