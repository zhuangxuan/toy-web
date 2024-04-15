package main

import (
	"fmt"
	"os"
)

func main() {
	// openfile可以打开目录，open不行,返回的不是一个文件指针，而是一个文件描述符

	// 获取用户输入的目录
	fmt.Println("请输入目录:")
	var dir string
	fmt.Scan(&dir)

	// 打开目录
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		// 读取目录下的文件
		file, err := f.Readdir(-1)
		if err != nil {
			panic(err)
		}
		if len(file) == 0 {
			break
		}
		for _, fileInfo := range file {
			if fileInfo.IsDir() {
				fmt.Println("fileInfo,是一个目录", fileInfo.Name())
			} else {
				//if strings.HasSuffix(fileInfo.Name(), ".go") {
				//	fmt.Println("fileInfo,是一个文件", fileInfo.Name())
				//} else {
				//	fmt.Println("fileInfo,是一个文件ttt", fileInfo.Name())
				//}
				//找到目录下txt文件中read单词出现的次数
				fmt.Println("fileInfo,是一个文件", fileInfo.Name())
			}
		}
	}

}
