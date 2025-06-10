package main

import (
	"context"
	"fmt"
	"golang_learning/user_auth/backend/database" // Update import path
	"golang_learning/user_auth/backend/routes"   // Update import path
	"golang_learning/user_auth/backend/utils"    // Update import path
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc" // 引入正确的包
	"go.opentelemetry.io/otel/sdk/trace"
	// 引入 gRPC 支持
)

func initTracer() {
	// 创建一个导出器
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),                 // 使用 insecure 模式
		otlptracegrpc.WithEndpoint("localhost:4317"), // 设置 endpoint
	)
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		utils.Logger.Fatal("Failed to create exporter:", err)
	}

	// 创建一个 tracer provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
}

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		utils.Logger.Warn("No .env file found")
	}

	// 初始化日志
	utils.InitLogger()

	// 初始化数据库
	if err := database.InitDB(); err != nil {
		utils.Logger.Fatal("Failed to connect to database:", err)
	}

	// 初始化 Redis
	utils.InitRedis()

	// 初始化 OpenTelemetry
	initTracer()

	// 设置路由
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	// 启动服务器
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		utils.Logger.Fatal("Failed to start server:", err)
	}
}
