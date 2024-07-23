package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// 假设的用户存储，实际应用中应使用数据库
var users = map[string]string{
	"admin": "password", // 密码应该是加密存储的
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// 验证用户名和密码
		if validateCredentials(username, password) {
			// 登录成功，重定向到主页或其他页面
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Fprintln(w, "Invalid username or password")
		}
	} else {
		http.ServeFile(w, r, "login.html")
	}
}

func validateCredentials(username, password string) bool {
	storedPassword, exists := users[username]
	if !exists {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err == nil
}

func main() {
	http.HandleFunc("/login", loginHandler)

	// 假设主页需要登录验证
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 这里可以添加检查会话或令牌的代码
		fmt.Fprintln(w, "Welcome to the home page!")
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
