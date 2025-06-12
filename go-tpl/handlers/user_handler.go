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
    if r.Method == "GET" {
        tmpl := template.Must(template.ParseFiles(
            filepath.Join("templates", "layout.html"),
            filepath.Join("templates", "user_form.html"),
        ))
        
        data := PageData{
            Title: "用户信息提交",
            User:  &models.User{},
        }
        
        tmpl.Execute(w, data)
        return
    }
    
    if r.Method == "POST" {
        user := &models.User{
            Name:  r.FormValue("name"),
            Email: r.FormValue("email"),
            Phone: r.FormValue("phone"),
        }
        
        err := user.Save()
        if err != nil {
            tmpl := template.Must(template.ParseFiles(
                filepath.Join("templates", "layout.html"),
                filepath.Join("templates", "user_form.html"),
            ))
            
            data := PageData{
                Title:       "用户信息提交",
                Message:     "保存失败：" + err.Error(),
                MessageType: "error",
                User:        user,
            }
            
            tmpl.Execute(w, data)
            return
        }
        
        // 重定向到成功页面
        http.Redirect(w, r, "/user?success=true", http.StatusSeeOther)
    }
}