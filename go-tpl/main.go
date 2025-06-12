// main.go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "golang_learning/go-tpl/database"
    "golang_learning/go-tpl/handlers"
)

func main() {
    // 初始化数据库连接
    err := database.InitDB()
    if err != nil {
        log.Fatal("数据库连接失败:", err)
    }
    
    // 删除旧表（如果存在）
    _, err = database.DB.Exec(`DROP TABLE IF EXISTS users`)
    if err != nil {
        log.Fatal("删除旧表失败:", err)
    }
    
    _, err = database.DB.Exec(`DROP TABLE IF EXISTS blogs`)
    if err != nil {
        log.Fatal("删除博客表失败:", err)
    }
    
    // 创建用户表
    _, err = database.DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(50) NOT NULL UNIQUE,
            password VARCHAR(100) NOT NULL,
            name VARCHAR(100),
            email VARCHAR(100),
            phone VARCHAR(20),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        log.Fatal("创建用户表失败:", err)
    }

    // 创建博客表
    _, err = database.DB.Exec(`
        CREATE TABLE IF NOT EXISTS blogs (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(200) NOT NULL,
            content TEXT NOT NULL,
            summary VARCHAR(500),
            author VARCHAR(100) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        log.Fatal("创建博客表失败:", err)
    }
    
    // 设置静态文件服务
    fs := http.FileServer(http.Dir("templates/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // 设置路由
    http.HandleFunc("/", handlers.HomeHandler)  // 将首页设为首页
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/user", handlers.UserFormHandler)
    
    // 创建服务器
    srv := &http.Server{
        Addr: ":8080",
    }
    
    // 在独立的 goroutine 中启动服务器
    go func() {
        log.Println("服务器启动在 http://localhost:8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("监听失败: %s\n", err)
        }
    }()
    log.Println("服务器启动成功")
    
    // 等待中断信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("正在关闭服务器...")
    
    // 设置 5 秒的超时时间
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // 优雅关闭服务器
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("服务器关闭:", err)
    }
    
    log.Println("服务器已关闭")
}