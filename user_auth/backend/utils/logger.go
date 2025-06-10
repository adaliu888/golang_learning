package utils

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Logger = logrus.New()

// InitLogger 初始化日志
func InitLogger() {
    Logger.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
    Logger.SetLevel(logrus.InfoLevel)

    // 将日志输出到文件
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        Logger.SetOutput(file)
    } else {
        Logger.Info("Failed to log to file, using default stderr")
    }
}