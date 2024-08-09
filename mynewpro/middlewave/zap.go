package middlewave

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*type config struct {
	Filename   string `mapstructure:"filename" json:"filename"`       // 日志文件路径
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`       // 日志文件最大大小（MB）
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"` // 保留旧文件的最大个数
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`         // 保留旧文件的最大天数
}

// viper获取配置信息
func Getconfiginfo() (config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}
*/

func InitLogger() *zap.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "/logs/myapp.log", // 日志文件路径
		MaxSize:    100,               // 日志文件最大大小（MB）
		MaxBackups: 5,                 // 保留旧文件的最大个数
		MaxAge:     28,                // 保留旧文件的最大天数
		Compress:   true,              // 是否压缩旧文件
	}

	writerSyncer := zapcore.AddSync(lumberJackLogger) // 将 lumberjack 包装为 zap 可以使用的格式

	encoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "ts",
		CallerKey:  "caller",
		EncodeTime: zapcore.ISO8601TimeEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writerSyncer,
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// GinLogger 使用 Zap 记录访问日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		zap.L().Info("access log",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery 使用 Zap 记录 panic 信息
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic error",
					zap.Any("err", err),
					zap.String("request", c.Request.URL.Path),
					zap.Stack("stack"),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
