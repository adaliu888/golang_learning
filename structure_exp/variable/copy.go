package copy

import (
	"fmt"
)

func QCopy() {
	// 定义一个切片
	slice := []int{1, 2, 3, 4, 5}

	// 复制切片到新切片
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	// 打印新切片
	fmt.Println(newSlice)

	// 改变新切片的值
	newSlice[0] = 100
	fmt.Println(newSlice)

	// 打印原始切片
	fmt.Println(slice)
}

func QCopy2() {
	type MyStruct struct {
		Field1 int
		Field2 []int
	}
	//原始数据
	original := MyStruct{
		Field1: 1,
		Field2: []int{1, 2, 3},
	}

	// 浅拷贝
	shallowCopy := original
	shallowCopy.Field2[0] = 100
	fmt.Println("Original after shallow copy modification:", original.Field2) // 输出：[100 2 3]

	// 深拷贝,change not modify original data
	var deepCopy MyStruct

	deepCopy.Field2 = make([]int, len(original.Field2))
	copy(deepCopy.Field2, original.Field2)
	deepCopy.Field2[0] = 200
	fmt.Println("Original after deep copy modification:", original.Field2) // 输出：[100 2 3]
	fmt.Println("Deep copy:", deepCopy.Field2)                             // 输出：[200 2 3]
}
