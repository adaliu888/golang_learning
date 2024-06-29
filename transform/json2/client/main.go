package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request 结构体用于定义客户端发送的JSON请求的结构
type Request struct {
	Message string `json:"message"`
}

func main() {
	// 创建一个Request实例
	req := Request{Message: "Hi there!"}

	// 将Request实例编码为JSON
	jsonBody, _ := json.Marshal(req)
	fmt.Printf("JSON request body: %s\n", jsonBody)

	// 发送HTTP POST请求
	resp, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	fmt.Println("Response Body:", resp.Body)
	fmt.Println("Response Body:", resp.ContentLength)
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 解码JSON响应
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	// 打印响应
	fmt.Printf("Received JSON response: %+v\n", response)
}
