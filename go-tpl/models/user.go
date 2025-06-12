// models/user.go
package models

import (
    "golang_learning/go-tpl/database"
)

type User struct {
    ID        int64
    Username  string
    Name      string
    Email     string
    Phone     string
    CreatedAt string
}

func (u *User) Save() error {
    // 更新用户信息
    query := `UPDATE users SET name = ?, email = ?, phone = ? WHERE username = ?`
    _, err := database.DB.Exec(query, u.Name, u.Email, u.Phone, u.Username)
    if err != nil {
        return err
    }
    return nil
}

func (u *User) GetByUsername(username string) error {
    query := `SELECT id, username, name, email, phone, created_at FROM users WHERE username = ?`
    err := database.DB.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Name, &u.Email, &u.Phone, &u.CreatedAt)
    if err != nil {
        return err
    }
    return nil
}