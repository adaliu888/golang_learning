package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 获取当前工作目录

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return ""
	}
	return dir
}

// 获取所有子目录

func GetSubDirectories(dir string) ([]string, error) {
	fileinfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var subDirectories []string
	for _, fileInfo := range fileinfos {
		if fileInfo.IsDir() {
			subDirectories = append(subDirectories, fileInfo.Name())
		}
	}

	return subDirectories, nil
}
