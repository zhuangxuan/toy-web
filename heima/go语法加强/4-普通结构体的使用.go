package main

import "fmt"

// 接口体内部字段定义顺序必须一致才能赋值

type Persion struct {
	name string
	sex  byte
	age  int
}

func main() {
	var temp Person

	fmt.Printf("&temp=%p/n", &temp)
	fmt.Printf("&temp.name=%p/n", &temp.name)
	//&temp=0x1400012c000/n&temp.name=0x1400012c000/n
	//	结构体的内存地址和结构体首个元素的地址是一样的

}
