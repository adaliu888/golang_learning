package file

import (
	"log"
	"os"
	"time"
)

// 编写一个函数，在指定路径下创建并记录当前日期的日志

func FN() (fn string) {
	date := time.Now().Format("2006-01-02")
	logFilename := date + ".log"

	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建一个日志记录器，将日志写入文件
	logger := log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// 记录一条日志信息
	logger.Println("这是一条带日期的日志信息")
	return logFilename
}
