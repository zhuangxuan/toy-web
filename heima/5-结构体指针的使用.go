package main

import (
	"fmt"
	"strings"
)

// 指针变量定义和初始化
// 1. 顺序初始化 依次将所有内部成员变量初始化

//func main() {
//	var p1 *Person = &Person{
//		name: "sss",
//		sex:  'a',
//		age:  0,
//	}
//	fmt.Println(p1)
//
//	var p2 *Person = new(Person)
//	p2.name = "sss"
//	p2.sex = 'a'
//	p2.age = 0
//	fmt.Println(p2)
//}

type Person struct {
	name string
	sex  byte
	age  int
}

func test(p *Person) {
	p.name = "test"
	p.sex = 'b'
	p.age = 1
}

func main() {
	var p1 *Person = &Person{
		name: "sss",
		sex:  'a',
		age:  0,
	}
	test(p1)
	fmt.Println(p1)
	fmt.Println("p1.name = ", p1.name)
	// 输出字符sex
	fmt.Println("p1.sex = ", p1.sex)

	fmt.Println(strings.Contains("hello", "o"))
}

// 函数调用结束后,栈真释放，局部变量的地址不变，不再受系统保护，随时间可能分配给其他程序
