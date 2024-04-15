package main

import (
	"fmt"
	"time"
)

func main() {
	GoRoutine()
	time.Sleep(1 * time.Second)
}

func GoRoutine() {
	go func() {
		fmt.Println("I am a goroutine")
		//time.Sleep(1 * time.Second)
		fmt.Println("I am a goroutine, too")
	}()
	// 这里直接输出，不会等待十秒
	fmt.Println("I am here")
}
