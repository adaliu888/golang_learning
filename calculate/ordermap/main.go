package main

import (
	"fmt"
)

// 假设我们定义了一个命令映射
var cmds = map[any]func(map[any]any) any{
	"add": func(args map[any]any) any {
		// 假设我们期望 args 包含两个 int 类型的键 "a" 和 "b"
		a := args["a"].(int)
		b := args["b"].(int)
		return a + b
	},
	// 可以添加更多的命令和它们对应的函数
}

func main() {
	// 调用 "add" 命令
	args := map[any]any{
		"a": 5,
		"b": 10,
	}
	result := cmds["add"](args)
	fmt.Println("Result:", result) // 输出：Result: 15
}
