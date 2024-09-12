package main

import (
	"fmt"
	"sort"
	"strings"
)

func OperaterStr() {
	var str string = "hello world, golang,my name is ada"
	var substr string = "hello world, golang,my name is ada"

	pos1 := sort.Search(len(str), func(i int) bool { return str[i] >= substr[0] })
	fmt.Println(pos1)

	pos := strings.Contains(str, substr) //判断str中是否包含substr
	fmt.Println(pos)
	//检查子字符串substr在字符串str中第一次出现的位置
	pos2 := strings.Index(str, substr)
	if pos2 == -1 {
		fmt.Println("not found")
	} else {
		fmt.Println(pos2)
	}

	//检查字符串最后出现的位置
	pos3 := strings.LastIndex(str, substr)
	fmt.Println(pos3)
	//分割字符串
	substrs := strings.Split(substr, ", ")
	fmt.Println(substrs)
	// 合并字符串
	joined := strings.Join(substrs, " and ")
	fmt.Println("Joined:", joined)

	// 替换子字符串
	replaced := strings.Replace(str, "world", "gopher", 1)
	fmt.Println("Replaced:", replaced)

}
