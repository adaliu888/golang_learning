package main

import (
	"fmt"
	"reflect"
)

// nil 是 Go 语言中的预声明标识符，代表指针、通道、函数、接口、map 或切片类型的零值

// 1. 不同类型的 nil 值示例
func nilTypesExample() {
	fmt.Println("\n=== 不同类型的 nil 值示例 ===")
	
	// 指针类型的 nil
	var p *int
	fmt.Printf("指针类型的 nil: %v, 类型: %T\n", p, p)
	
	// 通道类型的 nil
	var ch chan int
	fmt.Printf("通道类型的 nil: %v, 类型: %T\n", ch, ch)
	
	// 函数类型的 nil
	var f func()
	fmt.Printf("函数类型的 nil: %v, 类型: %T\n", f, f)
	
	// 接口类型的 nil
	var i interface{}
	fmt.Printf("接口类型的 nil: %v, 类型: %T\n", i, i)
	
	// map 类型的 nil
	var m map[string]int
	fmt.Printf("map 类型的 nil: %v, 类型: %T\n", m, m)
	
	// 切片类型的 nil
	var s []int
	fmt.Printf("切片类型的 nil: %v, 类型: %T\n", s, s)
}

// 2. nil 值的比较示例
func nilComparisonExample() {
	fmt.Println("\n=== nil 值的比较示例 ===")
	
	// 相同类型的 nil 比较
	var p1 *int
	var p2 *int
	fmt.Printf("相同类型的 nil 比较: %v\n", p1 == p2)
	
	// 不同类型的 nil 比较
	var p *int
	var ch chan int
	fmt.Printf("不同类型的 nil 比较: %v\n", p == nil && ch == nil)
	
	// 接口类型的 nil 比较
	var i1 interface{}
	var i2 interface{}
	fmt.Printf("接口类型的 nil 比较: %v\n", i1 == i2)
	
	// 接口与具体类型的 nil 比较
	var i interface{} = (*int)(nil)
	fmt.Printf("接口与具体类型的 nil 比较: %v\n", i == nil)
}

// 3. nil 值的使用示例
func nilUsageExample() {
	fmt.Println("\n=== nil 值的使用示例 ===")
	
	// 指针类型的 nil
	var p *int
	if p == nil {
		fmt.Println("指针是 nil")
	}
	
	// 通道类型的 nil
	var ch chan int
	if ch == nil {
		fmt.Println("通道是 nil")
	}
	// 注意：向 nil 通道发送或接收数据会导致永久阻塞
	
	// 函数类型的 nil
	var f func()
	if f == nil {
		fmt.Println("函数是 nil")
	}
	// 注意：调用 nil 函数会导致 panic
	
	// 接口类型的 nil
	var i interface{}
	if i == nil {
		fmt.Println("接口是 nil")
	}
	
	// map 类型的 nil
	var m map[string]int
	if m == nil {
		fmt.Println("map 是 nil")
	}
	// 注意：向 nil map 写入会导致 panic
	
	// 切片类型的 nil
	var s []int
	if s == nil {
		fmt.Println("切片是 nil")
	}
	// nil 切片可以安全地使用 len 和 cap
	fmt.Printf("nil 切片的长度: %d, 容量: %d\n", len(s), cap(s))
}

// 4. nil 接口的特殊性示例
func nilInterfaceExample() {
	fmt.Println("\n=== nil 接口的特殊性示例 ===")
	
	// 空接口的 nil
	var i1 interface{}
	fmt.Printf("空接口的 nil: %v, 类型: %T\n", i1, i1)
	
	// 包含 nil 指针的接口
	var p *int
	var i2 interface{} = p
	fmt.Printf("包含 nil 指针的接口: %v, 类型: %T\n", i2, i2)
	fmt.Printf("i2 == nil: %v\n", i2 == nil)
	
	// 使用反射检查 nil
	fmt.Printf("使用反射检查 i1 是否为 nil: %v\n", reflect.ValueOf(i1).IsNil())
	fmt.Printf("使用反射检查 i2 是否为 nil: %v\n", reflect.ValueOf(i2).IsNil())
}

// 5. nil 值的安全使用示例
func nilSafetyExample() {
	fmt.Println("\n=== nil 值的安全使用示例 ===")
	
	// 安全的指针使用
	var p *int
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("指针是 nil，安全跳过")
	}
	
	// 安全的 map 使用
	var m map[string]int
	if m != nil {
		m["key"] = 1
	} else {
		fmt.Println("map 是 nil，需要初始化")
		m = make(map[string]int)
		m["key"] = 1
	}
	
	// 安全的切片使用
	var s []int
	if s != nil {
		s = append(s, 1)
	} else {
		fmt.Println("切片是 nil，可以安全地 append")
		s = append(s, 1)
	}
	
	// 安全的通道使用
	var ch chan int
	if ch != nil {
		ch <- 1
	} else {
		fmt.Println("通道是 nil，需要初始化")
		ch = make(chan int)
		go func() {
			ch <- 1
		}()
	}
}

// 6. nil 值的常见错误示例
func nilErrorExample() {
	fmt.Println("\n=== nil 值的常见错误示例 ===")
	
	// 1. 解引用 nil 指针
	var p *int
	// fmt.Println(*p)  // 这会导致 panic
	
	// 2. 向 nil map 写入
	var m map[string]int
	// m["key"] = 1  // 这会导致 panic
	
	// 3. 调用 nil 函数
	var f func()
	// f()  // 这会导致 panic
	
	// 4. 向 nil 通道发送数据
	var ch chan int
	// ch <- 1  // 这会导致永久阻塞
	
	// 5. 从 nil 通道接收数据
	// <-ch  // 这会导致永久阻塞
	
	// 6. 关闭 nil 通道
	// close(ch)  // 这会导致 panic
	
	fmt.Println("以上操作都被注释掉了，因为它们会导致错误")
}

func main() {
	// 运行所有示例
	nilTypesExample()
	nilComparisonExample()
	nilUsageExample()
	nilInterfaceExample()
	nilSafetyExample()
	nilErrorExample()
} 