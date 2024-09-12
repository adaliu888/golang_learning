package main

import (
	"fmt"
	"time"
)

func main() {
	//打印當前的日志
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("CURRENT TIME: ", time.Now().Year(), time.Now().Month(), time.Now().Day())
}
