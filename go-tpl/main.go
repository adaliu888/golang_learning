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
    
    // 创建用户表
    _, err = database.DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) NOT NULL,
            phone VARCHAR(20) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        log.Fatal("创建用户表失败:", err)
    }
    
    // 设置静态文件服务
    fs := http.FileServer(http.Dir("templates/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // 设置路由
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