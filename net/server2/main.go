package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

// 定义日志和追踪数据的存储路径
const (
	LogDir    = "logs"
	TraceDir  = "traces"
	LogFile   = "app.log"
	TraceFile = "traces.json"
)

// initLogger 初始化日志系统
func initLogger() error {
	// 创建日志目录
	if err := os.MkdirAll(LogDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// 创建日志文件
	logPath := filepath.Join(LogDir, LogFile)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	// 配置 logrus
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Info("Logger initialized successfully")
	return nil
}

// initTracer 初始化 OpenTelemetry 追踪器
func initTracer() error {
	// 创建追踪数据目录
	if err := os.MkdirAll(TraceDir, 0755); err != nil {
		return fmt.Errorf("failed to create trace directory: %v", err)
	}

	// 创建追踪数据文件
	tracePath := filepath.Join(TraceDir, TraceFile)
	traceFile, err := os.OpenFile(tracePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open trace file: %v", err)
	}

	// 创建导出器，配置为将追踪数据写入文件
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(traceFile),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
	if err != nil {
		return fmt.Errorf("failed to create exporter: %v", err)
	}

	// 创建一个 Tracer Provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)

	logrus.Info("Tracer initialized successfully")
	return nil
}

// apiHandler 处理 API 请求的函数
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// 记录请求开始
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
		"ip":     r.RemoteAddr,
	}).Info("Received API request")

	// 从请求中获取追踪信息
	_, span := otel.Tracer("example.com/api").Start(r.Context(), "apiHandler")
	defer span.End()

	// 记录请求信息到追踪
	span.SetAttributes(
		attribute.String("http.method", r.Method),
		attribute.String("http.url", r.URL.String()),
		attribute.String("http.client_ip", r.RemoteAddr),
	)

	// 模拟一些工作
	time.Sleep(1 * time.Second)

	// 记录响应信息
	span.AddEvent("Processing API request")
	span.SetAttributes(attribute.String("http.status", "200"))

	// 处理请求
	fmt.Fprintf(w, "API request processed")

	// 记录请求完成
	logrus.WithFields(logrus.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
		"status": "200",
	}).Info("API request processed successfully")
}

func main() {
	// 初始化日志系统
	if err := initLogger(); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}

	// 初始化 OpenTelemetry
	if err := initTracer(); err != nil {
		logrus.Fatal("Failed to initialize tracer:", err)
	}

	// 设置路由
	http.HandleFunc("/api", apiHandler)

	// 启动服务器
	logrus.Info("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatal("Failed to start server:", err)
	}
}
