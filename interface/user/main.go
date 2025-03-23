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

/*
代码说明
mux 包：
mux 是 Go 语言中一个流行的路由器库，通常用于处理 HTTP 请求的路由。它允许你定义 URL 路径与处理函数之间的映射关系。
NewRouter 方法：
mux.NewRouter() 是 mux 包中的一个函数，用于创建一个新的路由器实例。这个路由器将用于定义和管理 HTTP 路由。
变量 r：
r := 是 Go 语言中的短变量声明语法，表示声明一个变量 r 并初始化为 mux.NewRouter() 的返回值。r 将是一个路由器对象，后续可以使用这个对象来定义路由规则。
*/
