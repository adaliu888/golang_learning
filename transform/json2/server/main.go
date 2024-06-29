package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response 结构体用于定义JSON响应的结构
type Response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 创建一个Response实例
		res := Response{Message: "Hello, World!"}

		// 将Response实例编码为JSON
		jsonResponse, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 设置Content-Type为application/json
		w.Header().Set("Content-Type", "application/json")
		// 写入JSON响应
		w.Write(jsonResponse)
	})

	// 启动HTTP服务器
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
