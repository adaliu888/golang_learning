package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// 全局中间件，记录请求日志
	r.Use(middleware.Logger)

	// 定义一个处理函数
	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home page")
	}

	// 定义用户路由分组
	userRoutes := chi.NewRouter()
	userRoutes.Get("/", homeHandler) // 用户主页
	userRoutes.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User profile page")
	})
	userRoutes.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User login")
	})

	// 将用户路由分组挂载到主路由器的 "/users" 路径下
	r.Mount("/users", userRoutes)

	// 定义管理路由分组
	adminRoutes := chi.NewRouter()
	adminRoutes.Use(middleware.BasicAuth("admin", map[string]string{"admin": "password"}))
	adminRoutes.Get("/", homeHandler) // 管理主页
	adminRoutes.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Admin dashboard")
	})

	// 将管理路由分组挂载到主路由器的 "/admin" 路径下
	r.Mount("/admin", adminRoutes)

	// 启动服务器
	http.ListenAndServe(":8080", r)
}
