package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 定义响应结构体
type ResponseData struct {
	Message string `json:"message"`
}



func main() {
	http.HandleFunc("/get-message", func(w http.ResponseWriter, r *http.Request) {
		// 创建响应数据
		resp := ResponseData{Message: "Hello from Go!"}
		// 将响应数据序列化为 JSON
		respBytes, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 设置响应类型为 JSON
		w.Header().Set("Content-Type", "application/json")
		// 写入响应
		w.Write(respBytes)
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
