package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 假设我们有一个User结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// getUsersHandler 是处理GET请求的函数
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// 从查询字符串中解析参数
	usersPerPage := r.URL.Query().Get("usersPerPage")
	page := r.URL.Query().Get("page")

	fmt.Print(usersPerPage)
	fmt.Print(page)
	// 根据参数进行业务逻辑处理，这里只是示例，实际中可能涉及数据库查询等操作
	// 假设我们有以下用户数据
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		// ... 更多用户
	}

	// 将用户数据序列化为JSON格式的响应
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()

	// 定义路由，使用mux的路径变量功能
	r.HandleFunc("/users", getUsersHandler).Methods("GET")

	// 启动HTTP服务器
	http.ListenAndServe(":8080", r)
}
