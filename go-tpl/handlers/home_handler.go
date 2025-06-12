package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
    "golang_learning/go-tpl/models"
)

type HomePageData struct {
    Title string
    Blogs []*models.Blog
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        filepath.Join("templates", "layout.html"),
        filepath.Join("templates", "home.html"),
    ))

    blog := &models.Blog{}
    blogs, err := blog.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := HomePageData{
        Title: "首页",
        Blogs: blogs,
    }

    tmpl.Execute(w, data)
} 