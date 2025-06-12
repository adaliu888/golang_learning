package models

import (
    "time"
    "golang_learning/go-tpl/database"
)

type Blog struct {
    ID        int64
    Title     string
    Content   string
    Summary   string
    Author    string
    CreatedAt time.Time
}

func (b *Blog) GetAll() ([]*Blog, error) {
    rows, err := database.DB.Query(`
        SELECT id, title, content, summary, author, created_at 
        FROM blogs 
        ORDER BY created_at DESC
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var blogs []*Blog
    for rows.Next() {
        blog := &Blog{}
        err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Summary, &blog.Author, &blog.CreatedAt)
        if err != nil {
            return nil, err
        }
        blogs = append(blogs, blog)
    }
    return blogs, nil
} 