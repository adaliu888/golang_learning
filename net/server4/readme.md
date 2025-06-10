我将为您创建一个示例，展示如何在 Go 中使用 OpenTelemetry 的 context 实现跨级联结构的追踪。这个示例将包含多层级的服务调用，展示如何传递和关联追踪信息。

```go
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
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/trace"
)

// 定义日志和追踪数据的存储路径
const (
    LogDir     = "logs"
    TraceDir   = "traces"
    LogFile    = "app.log"
    TraceFile  = "traces.json"
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

    // 模拟处理时间
    time.Sleep(100 * time.Millisecond)

    // 记录处理结果
    result := fmt.Sprintf("Processed by %s: %s", s.name, data)
    span.SetAttributes(attribute.String("output.result", result))

    return result, nil
}

// 服务B的方法
func (s *Service) ProcessB(ctx context.Context, data string) (string, error) {
    // 创建新的span
    ctx, span := otel.Tracer(s.name).Start(ctx, "ProcessB")
    defer span.End()

    // 记录输入数据
    span.SetAttributes(attribute.String("input.data", data))

    // 模拟处理时间
    time.Sleep(150 * time.Millisecond)

    // 记录处理结果
    result := fmt.Sprintf("Processed by %s: %s", s.name, data)
    span.SetAttributes(attribute.String("output.result", result))

    return result, nil
}

// 服务C的方法
func (s *Service) ProcessC(ctx context.Context, data string) (string, error) {
    // 创建新的span
    ctx, span := otel.Tracer(s.name).Start(ctx, "ProcessC")
    defer span.End()

    // 记录输入数据
    span.SetAttributes(attribute.String("input.data", data))

    // 模拟处理时间
    time.Sleep(200 * time.Millisecond)

    // 记录处理结果
    result := fmt.Sprintf("Processed by %s: %s", s.name, data)
    span.SetAttributes(attribute.String("output.result", result))

    return result, nil
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
        span.SetStatus(trace.StatusCodeError, err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 调用服务B
    resultB, err := serviceB.ProcessB(ctx, resultA)
    if err != nil {
        span.RecordError(err)
        span.SetStatus(trace.StatusCodeError, err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 调用服务C
    resultC, err := serviceC.ProcessC(ctx, resultB)
    if err != nil {
        span.RecordError(err)
        span.SetStatus(trace.StatusCodeError, err.Error())
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
```

这个示例展示了如何实现跨级联结构的追踪：

1. **服务层级结构**：
   - ServiceA -> ServiceB -> ServiceC
   - 每个服务都有自己的 span
   - 通过 context 传递追踪信息

2. **追踪信息传递**：
   - 使用 context 在服务之间传递追踪信息
   - 每个服务创建新的 span，但保持与父 span 的关联
   - 记录输入输出数据和处理时间

3. **错误处理**：
   - 记录错误信息到 span
   - 设置适当的错误状态

4. **使用说明**：

运行应用程序：
```bash
go run main.go
```

测试 API：
```bash
curl http://localhost:8080/process
```

5. **追踪数据**：
   - 追踪数据将保存在 `traces/traces.json` 文件中
   - 可以看到完整的调用链和每个服务的处理时间

6. **日志输出**：
   - 应用日志保存在 `logs/app.log` 文件中
   - 包含详细的请求处理信息

这个示例展示了如何：
- 创建多层级服务
- 使用 context 传递追踪信息
- 记录详细的追踪数据
- 处理错误情况
- 关联不同服务的追踪信息

如果您需要任何调整或有其他问题，请告诉我！
