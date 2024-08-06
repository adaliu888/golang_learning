package main

import (
	"fmt"
)

func main() {
	s := new(int) //分配内存
	rust := *s
	fmt.Println(rust) //返回0
	fmt.Println(s)    //返回指针指向的地址

	slic := make(map[string]int) //make是map的内建函数

	fmt.Println(slic) //返回空map
}
