// models/login_user.go
package models

import (
    "golang_learning/go-tpl/database"
    "fmt"
)

type LoginUser struct {
    ID        int64
    Username  string
    Password  string
}

func (u *LoginUser) Verify() error {
    var storedPassword string
    err := database.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", u.Username).Scan(&u.ID, &storedPassword)
    if err != nil {
        return err
    }

    if u.Password != storedPassword {
        return fmt.Errorf("密码错误")
    }

    return nil
} 