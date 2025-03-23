package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// 定义一个结构体来表示响应数据
type ResponseData struct {
	Message string `json:"message"`
}

// 定义一个处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := ResponseData{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	// 定义一个路由
	r.HandleFunc("/hello", helloHandler).Methods("GET")

	// 启动服务器
	http.ListenAndServe(":8080", r)
}
