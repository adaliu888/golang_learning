//slog := zap 结构化日志库

package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*func main() {

	logger, _ := zap.NewProduction()

	sugar := logger.Sugar()

	sugar.Infow("failed to fetch URL", "URL", "https://example.com", "attempt", 3, "backoff", 100*time.Millisecond)
}*/

func main() {
	cfg := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel), // 设置日志级别
		Encoding: "json",                              // 设置日志编码格式
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",    // 日志消息的键名
			LevelKey:   "level",  // 日志级别的键名
			TimeKey:    "ts",     // 时间戳的键名
			NameKey:    "logger", // 日志名称的键名
			// 更多配置...
		},
		// OutputPaths: []string{"stdout"},                  // 输出到标准输出（或其他路径）
		// ErrorOutputPaths: []string{"stderr"},            // 错误输出路径
		// 更多配置...
	}

	logger, err := cfg.Build()
	if err != nil {
		// 处理错误
	}
	defer logger.Sync() // 确保所有日志都已写入

	// 使用logger记录日志
	logger.Info("Hello, Zap")
}

/*配置Encoder:
zap使用EncoderConfig来配置日志的编码方式和输出格式。您可以设置键名、时间格式、级别编码等。

设置日志级别:
通过AtomicLevel设置日志的级别，例如DebugLevel、InfoLevel、WarnLevel等。

配置输出:
您可以设置OutputPaths来决定日志输出到哪个文件或标准输出。

错误处理:
使用ErrorOutputPaths设置错误日志的输出路径。

高级配置 (可选):
zap还支持许多高级配置，例如添加日志字段、设置日志旋转、配置回调钩子等。

使用Logger:
创建完Logger后，您可以使用它来记录不同级别的日志。
*/
