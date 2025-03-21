package main

import (
	"go.uber.org/zap"
)

func main() {
	// 创建 Production 模式日志器（JSON 格式，适合生产环境）
	logger, err := zap.NewDevelopment()
	if err != nil {
		// 若初始化错误，则直接终止程序
		panic("无法初始化 zap 日志器: " + err.Error())
	}
	// 确保在程序退出前同步日志缓冲区，将缓存日志写入磁盘或终端
	defer logger.Sync()

	// 输出一条 Info 级别日志，并附加结构化字段
	logger.Info("启动服务成功",
		zap.String("service", "ExampleApp"),
		zap.Int("port", 8080),
	)

	// 示例：记录错误日志
	sampleErr := doSomething()
	if sampleErr != nil {
		logger.Error("处理请求失败", zap.Error(sampleErr))
	}
}

// 示例函数，返回一个错误
func doSomething() error {
	// 模拟一个错误场景
	return nil
}
