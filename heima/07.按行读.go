package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	go readFile("./test.txt", ch)
	// 等待协程执行完毕
	<-ch
}
func readFile(filename string, ch chan int) {
	// 按行读取文件
	// 1. 打开文件
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			ch <- 0
			break
		}
		fmt.Println(string(line))
	}
}
