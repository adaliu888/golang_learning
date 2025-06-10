package main

import (
	"fmt"
	"unsafe"
)

// 引用类型包括：slice、map、channel、interface、function、pointer

// 1. Slice（切片）示例
func sliceExample() {
	fmt.Println("\n=== Slice 引用类型示例 ===")
	
	// 创建切片
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1  // 引用同一个底层数组
	
	fmt.Printf("slice1: %v, 地址: %p\n", slice1, &slice1)
	fmt.Printf("slice2: %v, 地址: %p\n", slice2, &slice2)
	
	// 修改 slice2 会影响 slice1
	slice2[0] = 100
	fmt.Printf("修改后 slice1: %v\n", slice1)
	fmt.Printf("修改后 slice2: %v\n", slice2)
	
	// 切片扩容会创建新的底层数组
	slice2 = append(slice2, 6, 7, 8)
	fmt.Printf("扩容后 slice1: %v\n", slice1)
	fmt.Printf("扩容后 slice2: %v\n", slice2)
	
	// 切片的内存布局
	fmt.Printf("slice1 长度: %d, 容量: %d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1 底层数组地址: %p\n", &slice1[0])
}

// 2. Map（映射）示例
func mapExample() {
	fmt.Println("\n=== Map 引用类型示例 ===")
	
	// 创建 map
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	
	// 引用赋值
	map2 := map1  // 引用同一个 map
	
	fmt.Printf("map1: %v\n", map1)
	fmt.Printf("map2: %v\n", map2)
	
	// 修改 map2 会影响 map1
	map2["c"] = 3
	fmt.Printf("修改后 map1: %v\n", map1)
	fmt.Printf("修改后 map2: %v\n", map2)
	
	// map 的零值是 nil
	var nilMap map[string]int
	fmt.Printf("nil map: %v, 是否为 nil: %v\n", nilMap, nilMap == nil)
}

// 3. Channel（通道）示例
func channelExample() {
	fmt.Println("\n=== Channel 引用类型示例 ===")
	
	// 创建通道
	ch1 := make(chan int, 3)
	ch2 := ch1  // 引用同一个通道
	
	// 发送数据
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)  // 关闭通道会影响所有引用
	}()
	
	// 从 ch2 接收数据
	for v := range ch2 {
		fmt.Printf("从 ch2 接收: %d\n", v)
	}
	
	// 通道的零值是 nil
	var nilCh chan int
	fmt.Printf("nil channel: %v, 是否为 nil: %v\n", nilCh, nilCh == nil)
}

// 4. Interface（接口）示例
func interfaceExample() {
	fmt.Println("\n=== Interface 引用类型示例 ===")
	
	// 接口变量存储了类型信息和值
	var i interface{} = "hello"
	
	// 类型断言
	if s, ok := i.(string); ok {
		fmt.Printf("接口值: %v, 类型: %T\n", s, s)
	}
	
	// 接口的零值是 nil
	var nilInterface interface{}
	fmt.Printf("nil interface: %v, 是否为 nil: %v\n", nilInterface, nilInterface == nil)
	
	// 空接口可以存储任何类型
	var any interface{}
	any = 42
	fmt.Printf("any 存储整数: %v, 类型: %T\n", any, any)
	any = "hello"
	fmt.Printf("any 存储字符串: %v, 类型: %T\n", any, any)
}

// 5. Function（函数）示例
func functionExample() {
	fmt.Println("\n=== Function 引用类型示例 ===")
	
	// 函数类型
	type Handler func(int) int
	
	// 定义函数
	double := func(x int) int {
		return x * 2
	}
	
	// 函数引用
	handler := double
	fmt.Printf("handler(5): %d\n", handler(5))
	
	// 函数作为参数
	apply := func(f Handler, x int) int {
		return f(x)
	}
	fmt.Printf("apply(double, 5): %d\n", apply(double, 5))
}

// 6. Pointer（指针）示例
func pointerExample() {
	fmt.Println("\n=== Pointer 引用类型示例 ===")
	
	// 基本类型指针
	x := 42
	p1 := &x
	p2 := p1  // 引用同一个地址
	
	fmt.Printf("x 的值: %d, 地址: %p\n", x, &x)
	fmt.Printf("p1 指向的值: %d, 地址: %p\n", *p1, p1)
	fmt.Printf("p2 指向的值: %d, 地址: %p\n", *p2, p2)
	
	// 修改 p2 会影响 p1 指向的值
	*p2 = 100
	fmt.Printf("修改后 x 的值: %d\n", x)
	
	// 结构体指针
	type Person struct {
		Name string
		Age  int
	}
	
	person := Person{"张三", 30}
	pp1 := &person
	pp2 := pp1  // 引用同一个结构体
	
	// 修改 pp2 会影响 pp1 指向的结构体
	pp2.Age = 31
	fmt.Printf("修改后 person: %+v\n", person)
	
	// 指针的零值是 nil
	var nilPtr *int
	fmt.Printf("nil pointer: %v, 是否为 nil: %v\n", nilPtr, nilPtr == nil)
}

// 7. 内存布局示例
func memoryLayoutExample() {
	fmt.Println("\n=== 引用类型内存布局示例 ===")
	
	// 切片内存布局
	slice := []int{1, 2, 3}
	fmt.Printf("切片大小: %d 字节\n", unsafe.Sizeof(slice))
	
	// map 内存布局
	m := make(map[string]int)
	fmt.Printf("map 大小: %d 字节\n", unsafe.Sizeof(m))
	
	// channel 内存布局
	ch := make(chan int)
	fmt.Printf("channel 大小: %d 字节\n", unsafe.Sizeof(ch))
	
	// interface 内存布局
	var i interface{} = "hello"
	fmt.Printf("interface 大小: %d 字节\n", unsafe.Sizeof(i))
	
	// 函数内存布局
	f := func() {}
	fmt.Printf("函数大小: %d 字节\n", unsafe.Sizeof(f))
	
	// 指针内存布局
	var p *int
	fmt.Printf("指针大小: %d 字节\n", unsafe.Sizeof(p))
}

func main() {
	// 运行所有示例
	sliceExample()
	mapExample()
	channelExample()
	interfaceExample()
	functionExample()
	pointerExample()
	memoryLayoutExample()
} 