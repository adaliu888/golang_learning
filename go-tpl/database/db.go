// database/db.go
package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
    var err error
    DB, err = sql.Open("mysql", "root:test1234@tcp(localhost:3306)/userdb?parseTime=true")
    if err != nil {
        return err
    }
    
    // 测试数据库连接
    err = DB.Ping()
    if err != nil {
        return err
    }
    
    return nil
}