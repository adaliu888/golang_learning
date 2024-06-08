package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// identity check for the file extension
	isNotexistsFile()
	fmt.Println(isNotexistsFile())
	fmt.Println("the file exist") //fmt.Println(existsFile())
	openFile()
}

func inFilename() string {
	var filename string
	fmt.Println("Please input the filename: ")
	fmt.Scanln(&filename)
	return filename
}

func indirectory() string {
	var directory string
	fmt.Println("Please input the directory:")
	fmt.Scanln(&directory)
	return directory
}

func linkpath() string {
	filename := inFilename()
	directory := indirectory()
	file := filepath.Join(directory, filename)
	return file
}

func isNotexistsFile() bool {
	// 解析文件路径
	file := linkpath()

	// 使用os.Stat检查文件
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在
			return false
		} else {
			// 发生其他类型的错误
			fmt.Println("Error checking file:", err)
			return false
		}

	} else {
		// 文件存在
		return true
	}
}

func openFile() {
	file := linkpath()
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println(f)
}
