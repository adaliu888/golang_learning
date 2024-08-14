package zzap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func IintZapLogger() *zap.Logger {
	// 创建一个开发环境的Logger配置
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel && lvl < zapcore.WarnLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 设置日志文件的滚动配置
	logRotateInfo := lumberjack.Logger{
		Filename:   "./zlogs/info.log", // 日志文件路径
		MaxSize:    10,                 // 每个日志文件的最大大小（MB）
		MaxBackups: 5,                  // 保留旧文件的最大数量
		MaxAge:     7,                  // 保留旧文件的最大天数
		Compress:   true,               // 是否压缩旧文件
	}
	logRotateError := lumberjack.Logger{
		Filename:   "./zlogs/error.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   true,
	}
	logRotateDebug := lumberjack.Logger{
		Filename:   "./zlogs/debug.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   true,
	}

	// 创建不同的io.Writer，每个级别的日志输出到不同的文件
	infoWriter := zapcore.AddSync(&logRotateInfo)
	errorWriter := zapcore.AddSync(&logRotateError)
	debugWriter := zapcore.AddSync(&logRotateDebug)

	// 编码器配置
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// 创建核心Logger
	cores := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), infoWriter, infoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), errorWriter, warnLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), debugWriter, debugLevel),
	)

	// 使用核心Logger创建最终的Logger
	logger := zap.New(cores)
	// 确保退出时日志被刷新

	return logger

	// 使用ReplaceGlobals设置全局Logger
	//defer zap.ReplaceGlobals(logger)

	// 记录日志
	//logger.Debug("This is a debug message", zap.String("key", "value"))
	//logger.Info("This is an info message", zap.Int("number", 123))
	//logger.Warn("This is a warning")
	//logger.Error("This is an error message", zap.Error(nil))

	// 现在可以在任何地方使用zap.L()来获取Logger并记录日志

	//	for i := 0; i < 10; i++ {

	//zap.L().Info("Hello from global logger")
	//zap.L().Error("Hello from global logger")
	//zap.S().Debug("HELOG FOR GLOBAL LOGGER")
	//}

}
