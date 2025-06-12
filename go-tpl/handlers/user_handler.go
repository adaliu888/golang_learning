// handlers/user_handler.go
package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
    "golang_learning/go-tpl/models"
)

type PageData struct {
    Title       string
    Message     string
    MessageType string
    User        *models.User
}

func UserFormHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        filepath.Join("templates", "layout.html"),
        filepath.Join("templates", "user_form.html"),
    ))

    if r.Method == "GET" {
        // 从 session 或 cookie 中获取用户名
        // 这里暂时使用一个固定的用户名进行测试
        username := "test_user" // 实际应用中应该从 session 获取

        user := &models.User{Username: username}
        err := user.GetByUsername(username)
        if err != nil {
            // 如果用户不存在，创建一个新的用户对象
            user = &models.User{Username: username}
        }

        data := PageData{
            Title: "用户信息提交",
            User:  user,
        }
        
        tmpl.Execute(w, data)
        return
    }
    
    if r.Method == "POST" {
        // 从 session 或 cookie 中获取用户名
        // 这里暂时使用一个固定的用户名进行测试
        username := "test_user" // 实际应用中应该从 session 获取

        user := &models.User{
            Username: username,
            Name:     r.FormValue("name"),
            Email:    r.FormValue("email"),
            Phone:    r.FormValue("phone"),
        }
        
        err := user.Save()
        if err != nil {
            data := PageData{
                Title:       "用户信息提交",
                Message:     "提交失败",
                MessageType: "error",
                User:        user,
            }
            
            tmpl.Execute(w, data)
            return
        }
        
        // 提交成功后重定向到GET请求，重置表单
        http.Redirect(w, r, "/user", http.StatusSeeOther)
    }
}