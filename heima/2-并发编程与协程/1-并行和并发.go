package main

import (
	"fmt"
	"time"
)

// 定义channel

// make(chan Type) 等价于 make(chan Type, 0)
// make(chan Type, capacity)

// 定义一台打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c \n", ch)
		time.Sleep(3 * time.Second)
	}
}

// 每当有一个进程启动时，系统会自动打开三个文件：标准输入、标准输出和标准错误。---三个文件为：stdin、stdout、stderr
var ch = make(chan int)

func person1() {
	printer("hello")
	ch <- 1
}

func person2() {
	<-ch
	printer("world")
}

func main() {
	//NumCPU
	//fmt.Println(runtime.NumCPU())
	//fmt.Println(runtime.Version())

	go person1()
	go person2()

	for {

	}
}
