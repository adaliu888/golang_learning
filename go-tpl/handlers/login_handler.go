// handlers/login_handler.go
package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
    "golang_learning/go-tpl/models"
)

type LoginData struct {
    Title       string
    Message     string
    MessageType string
    User        *models.LoginUser
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        filepath.Join("templates", "layout.html"),
        filepath.Join("templates", "login.html"),
    ))

    if r.Method == "GET" {
        data := LoginData{
            Title: "用户登录",
            User:  &models.LoginUser{},
        }
        
        tmpl.Execute(w, data)
        return
    }
    
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // 验证用户
        user := &models.LoginUser{
            Username: username,
            Password: password,
        }
        
        err := user.Verify()
        if err != nil {
            data := LoginData{
                Title:       "用户登录",
                Message:     "用户名或密码错误",
                MessageType: "error",
                User:        user,
            }
            tmpl.Execute(w, data)
            return
        }
        
        // 登录成功后重定向到用户信息提交页面
        http.Redirect(w, r, "/user", http.StatusSeeOther)
    }
} 