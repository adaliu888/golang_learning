package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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

// 定义服务结构
type Service struct {
	name string
}

// 创建新的服务实例
func NewService(name string) *Service {
	return &Service{name: name}
}

// 初始化日志系统
func initLogger() error {
	if err := os.MkdirAll(LogDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	logPath := filepath.Join(LogDir, LogFile)
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Info("Logger initialized successfully")
	return nil
}

// 初始化追踪器
func initTracer() error {
	if err := os.MkdirAll(TraceDir, 0755); err != nil {
		return fmt.Errorf("failed to create trace directory: %v", err)
	}

	tracePath := filepath.Join(TraceDir, TraceFile)
	traceFile, err := os.OpenFile(tracePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open trace file: %v", err)
	}

	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(traceFile),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
	if err != nil {
		return fmt.Errorf("failed to create exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)

	logrus.Info("Tracer initialized successfully")
	return nil
}

// 服务A的方法
func (s *Service) ProcessA(ctx context.Context, data string) (string, error) {
	// 创建新的span
	ctx, span := otel.Tracer(s.name).Start(ctx, "ProcessA")
	defer span.End()

	// 记录输入数据
	span.SetAttributes(attribute.String("input.data", data))

	// 使用 context 进行超时控制
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// 模拟处理时间
	select {
	case <-ctx.Done():
		span.SetStatus(codes.Error, "context cancelled")
		return "", ctx.Err()
	case <-time.After(100 * time.Millisecond):
		// 记录处理结果
		result := fmt.Sprintf("Processed by %s: %s", s.name, data)
		span.SetAttributes(attribute.String("output.result", result))
		return result, nil
	}
}

// 服务A的方法
func (s *Service) ProcessB(ctx context.Context, data string) (string, error) {
	// 创建新的span
	ctx, span := otel.Tracer(s.name).Start(ctx, "ProcessA")
	defer span.End()

	// 记录输入数据
	span.SetAttributes(attribute.String("input.data", data))

	// 使用 context 进行超时控制
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// 模拟处理时间
	select {
	case <-ctx.Done():
		span.SetStatus(codes.Error, "context cancelled")
		return "", ctx.Err()
	case <-time.After(100 * time.Millisecond):
		// 记录处理结果
		result := fmt.Sprintf("Processed by %s: %s", s.name, data)
		span.SetAttributes(attribute.String("output.result", result))
		return result, nil
	}
}

// 服务C的方法
func (s *Service) ProcessC(ctx context.Context, data string) (string, error) {
	// 创建新的span
	ctx, span := otel.Tracer(s.name).Start(ctx, "ProcessC")
	defer span.End()

	// 记录输入数据
	span.SetAttributes(attribute.String("input.data", data))

	// 使用 context 进行超时控制
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// 模拟处理时间
	select {
	case <-ctx.Done():
		span.SetStatus(codes.Error, "context cancelled")
		return "", ctx.Err()
	case <-time.After(200 * time.Millisecond):
		// 记录处理结果
		result := fmt.Sprintf("Processed by %s: %s", s.name, data)
		span.SetAttributes(attribute.String("output.result", result))
		return result, nil
	}
}

// 处理HTTP请求的函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 创建根span
	ctx, span := otel.Tracer("http").Start(r.Context(), "handleRequest")
	defer span.End()

	// 记录请求信息
	span.SetAttributes(
		attribute.String("http.method", r.Method),
		attribute.String("http.url", r.URL.String()),
	)

	// 创建服务实例
	serviceA := NewService("ServiceA")
	serviceB := NewService("ServiceB")
	serviceC := NewService("ServiceC")

	// 模拟数据流
	data := "Hello, OpenTelemetry!"

	// 调用服务A
	resultA, err := serviceA.ProcessA(ctx, data)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 调用服务B
	resultB, err := serviceB.ProcessB(ctx, resultA)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 调用服务C
	resultC, err := serviceC.ProcessC(ctx, resultB)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 记录最终结果
	span.SetAttributes(attribute.String("final.result", resultC))
	fmt.Fprintf(w, "Final result: %s\n", resultC)
}

func main() {
	// 初始化日志系统
	if err := initLogger(); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}

	// 初始化追踪器
	if err := initTracer(); err != nil {
		logrus.Fatal("Failed to initialize tracer:", err)
	}

	// 设置路由
	http.HandleFunc("/process", handleRequest)

	// 启动服务器
	logrus.Info("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatal("Failed to start server:", err)
	}
}
