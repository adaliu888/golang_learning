// handlers/register_handler.go
package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
    "golang_learning/go-tpl/models"
)

type RegisterData struct {
    Title       string
    Message     string
    MessageType string
    User        *models.RegisterUser
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        filepath.Join("templates", "layout.html"),
        filepath.Join("templates", "register.html"),
    ))

    if r.Method == "GET" {
        data := RegisterData{
            Title: "用户注册",
            User:  &models.RegisterUser{},
        }
        
        tmpl.Execute(w, data)
        return
    }
    
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")
        confirmPassword := r.FormValue("confirm_password")

        // 验证用户名长度
        if len(username) < 4 || len(username) > 20 {
            data := RegisterData{
                Title:       "用户注册",
                Message:     "用户名长度必须在4-20个字符之间",
                MessageType: "error",
                User: &models.RegisterUser{
                    Username: username,
                },
            }
            tmpl.Execute(w, data)
            return
        }

        // 验证密码长度
        if len(password) < 6 || len(password) > 20 {
            data := RegisterData{
                Title:       "用户注册",
                Message:     "密码长度必须在6-20个字符之间",
                MessageType: "error",
                User: &models.RegisterUser{
                    Username: username,
                },
            }
            tmpl.Execute(w, data)
            return
        }

        // 验证两次密码是否一致
        if password != confirmPassword {
            data := RegisterData{
                Title:       "用户注册",
                Message:     "两次输入的密码不一致",
                MessageType: "error",
                User: &models.RegisterUser{
                    Username: username,
                },
            }
            tmpl.Execute(w, data)
            return
        }

        // 创建用户
        user := &models.RegisterUser{
            Username: username,
            Password: password,
        }
        
        err := user.Save()
        if err != nil {
            data := RegisterData{
                Title:       "用户注册",
                Message:     "注册失败：" + err.Error(),
                MessageType: "error",
                User:        user,
            }
            tmpl.Execute(w, data)
            return
        }
        
        // 注册成功后重定向到用户信息提交页面
        http.Redirect(w, r, "/user", http.StatusSeeOther)
    }
} 