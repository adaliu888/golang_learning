// models/user.go
package models

import (
    "golang_learning/go-tpl/database"
)

type User struct {
    ID        int64
    Name      string
    Email     string
    Phone     string
    CreatedAt string
}

func (u *User) Save() error {
    query := `INSERT INTO users (name, email, phone, created_at) VALUES (?, ?, ?, NOW())`
    result, err := database.DB.Exec(query, u.Name, u.Email, u.Phone)
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