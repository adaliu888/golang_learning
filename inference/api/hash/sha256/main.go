package main

import (
	"crypto/sha256" //hash sha256 支持的库(SHA1,MD5)
	"fmt"
	"io"
)

func main() {
	data := "Hello, World!" //data
	hash := sha256.New()    //sha256.New() 创建了一个新的 SHA256 哈希实例(sha1,DM5)

	// 向哈希状态机写入数据
	_, err := io.WriteString(hash, data) //使用 io.WriteString 向它写入了要哈希的数据
	if err != nil {
		panic(err)
	}

	// 获取哈希值
	hashed := hash.Sum(nil) //通过sum()方法获取哈希值

	fmt.Printf("SHA256 hash of '%s': %x\n", data, hashed)

}

//SHA256 hash of 'Hello, World!': dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f
