package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// 使用中间件记录请求日志
	r.Use(middleware.Logger)

	// 定义路由
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home page")
	})

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// 使用 Chi 的 URL 参数绑定
		userID := chi.URLParam(r, "id")
		fmt.Fprintf(w, "User with ID: %s", userID)
	})

	// 路由分组示例
	r.Route("/admin", func(r chi.Router) {
		r.Use(middleware.BasicAuth("admin", map[string]string{"admin": "password"}))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Admin home page")
		})
	})

	// 使用 Chi 启动服务器
	http.ListenAndServe(":8080", r)
}
