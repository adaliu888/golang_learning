// models/register_user.go
package models

import (
    "golang_learning/go-tpl/database"
    "fmt"
)

type RegisterUser struct {
    ID        int64
    Username  string
    Password  string
    CreatedAt string
}

func (u *RegisterUser) Save() error {
    // 检查用户名是否已存在
    var count int
    err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", u.Username).Scan(&count)
    if err != nil {
        return err
    }
    if count > 0 {
        return fmt.Errorf("用户名已存在")
    }

    // 插入新用户
    query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, NOW())`
    result, err := database.DB.Exec(query, u.Username, u.Password)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    u.ID = id
    return nil
} 