go get "go.opentelemetry.io/otel" \
  "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric" \
  "go.opentelemetry.io/otel/exporters/stdout/stdouttrace" \
  "go.opentelemetry.io/otel/propagation" \
  "go.opentelemetry.io/otel/sdk/metric" \
  "go.opentelemetry.io/otel/sdk/resource" \
  "go.opentelemetry.io/otel/sdk/trace" \
  "go.opentelemetry.io/otel/semconv/v1.24.0" \
  "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

应用执行
go mod tidy
export OTEL_RESOURCE_ATTRIBUTES="service.name=dice,service.version=0.1.0"
go run .

使用jeager
如果觉着控制台看的 span 不够直观，可以选择将链路追踪的数据发送至 Jaeger，通过 Jaeger UI 查看。

启动 Jaeger
Jaeger 官方提供的 all-in-one 是为快速本地测试而设计的可执行文件。它包括 Jaeger UI、jaeger-collector、jaeger-query 和 jaeger-agent，以及一个内存存储组件。

启动 all-in-one 的最简单方法是使用发布到 DockerHub 的预置镜像（只需一条命令行）。
