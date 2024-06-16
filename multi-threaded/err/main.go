//在go 语言中错误处理是编程的一个重要部分，因为go是一种静态类型，编译型语言，它不会在运行是自动处理类型错误或访问越界等问题

// 1、返回值错误
package main

import (
	"errors"
	"fmt"
	"os"
)

func Divide(x, y int) (result int, err error) {
	if y == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}
	return x / y, nil
}

//使用defer进行资源清理：确保函数返回之前执行一些操作，比如关闭文件或者释放资源

func Filedeal(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
}

// 2、检查错误
func main() {
	result, err := Divide(100, 0) //被除数为0，导致错误
	if err != nil {
		fmt.Println(err)
		myfunc()
		return

	}
	fmt.Println(result)
	//使用defer进行资源清理：
	Filedeal("a.txt")
	myfunc()

	// 3、panic and recover
	myfunc()
	anotherFunc()
	//5、错误包装，使用errors.Is和errors.As方法来检查和包装错误
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file not exist")
	}
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err)
	}
	if errors.As(err, &os.PathError{}) {
		fmt.Println(err)
	}
}

//4、使用panic and recover 处理异常：panic 可以在运行时终端程序的执行流，而recover可以捕获这个终端，恢复程序的运行。

func myfunc() {
	panic("myfunc panic problem")
}

func anotherFunc() {
	defer func() {
		if err := recover(); err != nil { //捕获异常,并恢复程序的运行
			fmt.Println("recover from ", err)
		}
	}()
	myfunc()
}
