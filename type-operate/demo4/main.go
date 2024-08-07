//在 Go 语言中，变长参数（也称为可变参数）允许你向函数传递任意数量的参数值。变长参数在函数定义时使用省略号 ... 表示，并且在函数体内部作为一个切片处理。

//以下是使用 Go 语言变长参数的一些关键点：

//定义变长参数：
//在函数的参数列表中使用 ... 来定义变长参数。例如：func myFunc(arg1 int, arg2 ...string)。

//类型一致性：
//变长参数可以是任何类型，但所有传递给变长参数的参数必须可以被转换成该类型。

//在函数内部使用：
//在函数内部，变长参数作为一个切片处理，可以使用切片的所有操作。

//遍历变长参数：
//你可以使用循环来遍历变长参数中的所有参数。

//变长参数的零值：
//如果调用函数时没有传递任何参数给变长参数，那么变长参数的切片长度为零。

//变长参数与类型检查：
//在使用变长参数时，可能需要进行类型检查和断言，尤其是在函数期望不同类型的参数时。

//变长参数的限制：
//变长参数在函数的参数列表中只能出现一次，并且必须是最后一个参数。

//以下是一个使用变长参数的示例：

package main

import "fmt"

// 定义一个接受变长参数的函数
func printValues(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	// 调用函数，传递不同数量和类型的参数
	printValues(10, "hello", true) // 输出: 10, hello, true
	printValues("single value")
	printValues() // 不传递任何参数
}
