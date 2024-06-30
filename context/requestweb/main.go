//访问网页

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var url = "www.baidu.com"

	// 根据操作系统使用不同的命令
	var cmd *exec.Cmd

	// Windows系统
	cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)

	// macOS系统
	// cmd = exec.Command("open", url)

	// Linux系统
	// cmd = exec.Command("xdg-open", url)

	// 启动浏览器
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Browser should open and navigate to %s\n", url)
	}

	// 等待命令执行完成（非必须）
	cmd.Wait()

}
