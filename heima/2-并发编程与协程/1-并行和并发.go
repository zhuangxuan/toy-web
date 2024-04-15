package main

import (
	"fmt"
	"runtime"
)

// 定义channel
var ch chan int

// make(chan Type) 等价于 make(chan Type, 0)
// make(chan Type, capacity)
func main() {
	//NumCPU
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())

}
