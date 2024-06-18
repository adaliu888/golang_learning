package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//string variable
	var s string = "hello"   //声明变量并赋值
	w := "world"             //短声明
	fmt.Println(s + " " + w) //字符串连接
	fmt.Println(len(s))      //字符串长度len(s)

	fmt.Println("go" + "lang") //字符串连接
	fmt.Println("1+1 =", 1+1)  //字符串连接
	// string index
	char := s[1:] //获取指定位置的字符,使用索引
	fmt.Println(char)

	//string slice
	fmt.Println(s[0:5]) //获取指定范围的字符串

	//string convert
	fmt.Println(string(char))
	//string compare
	if "apple" > "banana" {
		fmt.Println("apple != banana")
	}

	//查找子串

	substring := strings.Index(s, "ll")
	fmt.Println(substring)

	//替换子串

	fmt.Println(strings.ReplaceAll(s, "l", "p"))

	//大小写转换
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))

	//字符串分割

	fmt.Println(strings.Split(s, "l"))
	fmt.Println(strings.Split(s, ""))

	//去除空白

	p := "  geikm  wibuk  "

	fmt.Println(strings.TrimSpace(p))

	//字符串格式化
	fmt.Printf("hello %s", "world") // %s, %d, %f, %b, %c, %v, %T

	//regex
	re := regexp.MustCompile(`[a-z]+`)

	fmt.Println(re.FindAllString(s, 2))

	if re.MatchString(s) {
		fmt.Println("match")
	}

}
