package main

import (
	"fmt"
	"sort"
)

type MyIntSlice []int

// 实现 sort.Interface 接口的 Len 方法
func (s MyIntSlice) Len() int {
	return len(s)
}

// 实现 sort.Interface 接口的 Less 方法,比较两个元素，并返回第一个元素应该排在第二个之前
func (s MyIntSlice) Less(i, j int) bool {
	return s[i] < s[j] // 升序排序
}

// // Swap 是 sort.Interface 的方法，交换切片中的两个元素
func (s MyIntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {

	var a = []int{2, 3, 1, 5, 4}
	//对整数切片排序
	sort.Ints(a)
	fmt.Println(a)
	//对字符串切片排序
	var s = []string{"c", "a", "b", "S", "A"}
	sort.Strings(s)
	fmt.Println(s)
	//对字符串切片排序
	m := []int{5, 3, 4, 1, 2}
	sort.Slice(m, func(i, j int) bool {
		return m[i] > m[j] //降序排序
	})
	fmt.Println(m)
	//通用排序接口

	var b = []int{2, 3, 1, 5, 4}
	//对整数切片排序
	sort.Sort(sort.IntSlice(b))
	fmt.Println(b)

	var c = []string{"c", "a", "b", "S", "A"}
	//对字符串切片排序
	sort.Sort(sort.StringSlice(c))
	fmt.Println(c)

	//// 创建并初始化一个整数切片
	IntSlice := MyIntSlice{3, 4, 2, 6, 7, 1, 6} //

	// 对切片进行排序
	sort.Sort(IntSlice)
	fmt.Println(IntSlice)
	//字符进行排序
	// 字符串切片
	strSlice := []string{"go", "golang", "sort", "package"}
	sort.Strings(strSlice)
	fmt.Println("Sorted strings:", strSlice)

	//// 自定义类型排序
	type MyStruct struct {
		Name string
		Age  int
	}
	mySlice := []MyStruct{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 22},
		{Name: "Dave", Age: 25},
	}
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].Age < mySlice[j].Age
	})
	fmt.Println("Sorted structs by age:", mySlice)

	// 逆序排序
	sort.Reverse(sort.StringSlice(strSlice))
	fmt.Println("Reverse sorted structs:", strSlice)

}
